#!/usr/bin/env bash

# Convert MKV to MP4
# Uses Docker FFMpeg, FFMpeg or AVConv.
#
# Usage:
# ./convert_video.sh
# ./convert_video.sh /path/to/videos
#
# Notes:
# Statistics: -stats, -nostats, -loglevel 0
# Skip Subtitles: -sn
# Copy Subtitles: -c:s mov_text
#
# Links:
# https://askubuntu.com/questions/396883/how-to-simply-convert-video-files-i-e-mkv-to-mp4
# https://andre.blue/blog/converting-avi-to-mp4-with-ffmpeg/

VIDEO_PATH=${1:-.}

if [ ! -d "$VIDEO_PATH" ]; then
    echo "Directory '$VIDEO_PATH' does not exist"
    exit 1
fi

convert_video() {
    local video="$1"
    local file=$(basename "$video")
    local extension="${file##*.}"

    if [ "$extension" == "mkv" ]; then
        local opts="-nostats -loglevel 0 -c:v copy -c:a copy -c:s mov_text -movflags +faststart"
    elif [ "$extension" == "avi" ]; then
        local opts="-nostats -loglevel 0 -c:a aac -b:a 128k -c:v libx264 -crf 23 -movflags +faststart"
    fi

    echo "Processing '$(basename "$video")'"

    if which docker > /dev/null; then
        local path=$(dirname "$video")
        local file=$(basename "$video")

        docker run --rm -v="$path:/tmp/workdir" -w="/tmp/workdir" jrottenberg/ffmpeg -i "${file}" $opts "${file%.*}.mp4"
    elif which ffmpeg > /dev/null; then
        ffmpeg -i "${video}" $opts "${video%.*}.mp4"
    elif which avconv > /dev/null; then
        avconv -i "${video}" $opts "${video%.*}.mp4"
    fi

    if [ "$?" -ne "0" ]; then
        echo "Failed to convert video '$video'"
        exit 1
    fi
}

find "$VIDEO_PATH" \
    -type f \
    -name "*.mkv" \
    -or \
    -name "*.avi" \
    -not -iwholename "*.AppleDouble*" \
    -not -iwholename "*._*" \
    | while read file; do convert_video "$file"; done
