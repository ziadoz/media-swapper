#!/usr/bin/env bash
# https://hub.docker.com/r/jrottenberg/ffmpeg/tags/

for video in *mkv; do
    docker run --rm -v="`pwd`:/tmp/workdir" -w="/tmp/workdir" jrottenberg/ffmpeg \
        -stats \
        -i "${video}" \
        -c:v copy \
        -c:a copy \
        -c:s mov_text \
        -movflags \
        +faststart \
        "${video%.*}.mp4"
done
