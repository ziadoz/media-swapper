#!/usr/bin/env bash

# Usage:
# ./convert_video.sh
# ./convert_video.sh /path/to/videos

VIDEO_PATH=${1:-.}

if [ ! -d "$VIDEO_PATH" ]; then
    echo "ERROR: Directory '${VIDEO_PATH}' does not exist"
    exit 1
fi

convert_video() {
    local video="$1"
    local file=$(basename "$video")
    local path=$(dirname "$video")
    local backup="$path/_Backups"

    echo "Processing '${video}'"

    # Notes:
    # Statistics: -stats, -nostats, -loglevel 0
    # Skip Subtitles: -sn
    # Copy Subtitles: -c:s mov_text

    docker run --rm -v="$path:/tmp/workdir" -w="/tmp/workdir" jrottenberg/ffmpeg \
        -i "${file}" \
        -nostats \
        -loglevel 0 \
        -c:v copy \
        -c:a copy \
        -c:s mov_text \
        -movflags \
        +faststart \
        "${file%.*}.mp4"

    if [ "$?" -ne "0" ]; then
        echo "ERROR: Failed to convert video '${video}'"
        exit 1
    fi

    mkdir -p "$backup"
    mv "$video" "$backup"

    echo "Processed '${video}'"
}

find "$VIDEO_PATH" -type f -name "*.mkv" -not -iwholename "*.AppleDouble*" -not -iwholename "*._*" | while read file; do convert_video "$file"; done
