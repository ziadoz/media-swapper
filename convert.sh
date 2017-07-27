#!/usr/bin/env bash

VIDEO_PATH=${1:-.}

if [ ! -d "$VIDEO_PATH" ]; then
    echo "Directory '${VIDEO_PATH}' does not exist"
    exit
fi

convert_video() {
    local video="$1"
    local file=$(basename "$video")
    local path=$(dirname "$video")
    local backup="$path/_Backups"

    # Notes:
    # Skip Subtitles: -sn
    # Copy Subtitles: -c:s mov_text

    docker run --rm -v="$path:/tmp/workdir" -w="/tmp/workdir" jrottenberg/ffmpeg \
        -stats \
        -i "${file}" \
        -c:v copy \
        -c:a copy \
        -c:s mov_text
        -movflags \
        +faststart \
        "${file%.*}.mp4"

    mkdir -p "$backup"
    mv "$video" "$backup"
}

find "$VIDEO_PATH" -type f -name "*.mkv" -not -iwholename "*.AppleDouble*" -not -iwholename "*._*" | while read file; do convert_video "$file"; done
