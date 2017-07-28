#!/usr/bin/env bash

# Usage:
# ./convert_video.sh
# ./convert_video.sh /path/to/videos

VIDEO_PATH=${1:-.}
BACKUP_DIR="_Backups"

if [ ! -d "$VIDEO_PATH" ]; then
    echo "Directory '${VIDEO_PATH}' does not exist"
    exit 1
fi

convert_video() {
    local video="$1"
    local path=$(dirname "$video")
    local file=$(basename "$video")
    local backup="$path/$BACKUP_DIR"

    if [ -e "${video%.*}.mp4" ]; then
        rm -rf "${video%.*}.mp4"
    fi

    echo "Processing '$video'"

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
        echo "Failed to convert video '${video}'"
        exit 1
    fi

    mkdir -p "$backup"
    mv "$video" "$backup"

    echo "Processed '$video'"
}

find "$VIDEO_PATH" \
    -type f \
    -name "*.mkv" \
    -not -iwholename "*.AppleDouble*" \
    -not -iwholename "*._*" \
    -not -wholename "*$BACKUP_DIR*" \
    | while read file; do convert_video "$file"; done
