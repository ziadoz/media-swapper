#!/usr/bin/env bash

# Links
# https://hub.docker.com/r/jrottenberg/ffmpeg/tags/
# https://gist.github.com/pwenzel/d6a0b54b120afac0bd1f
# https://askubuntu.com/questions/396883/how-to-simply-convert-video-files-i-e-mkv-to-mp4
# http://www.bugcodemaster.com/article/convert-videos-mp4-format-using-ffmpeg
# https://askubuntu.com/questions/50433/how-to-convert-mkv-file-into-mp4-file-losslessly
# http://stackoverflow.com/questions/1224766/how-do-i-rename-the-extension-for-a-batch-of-files
# https://superuser.com/questions/525249/convert-avi-to-mp4-keeping-the-same-quality
# https://github.com/PHP-FFMpeg/PHP-FFMpeg
# https://hub.docker.com/r/jrottenberg/ffmpeg/tags/
#
# Examples
# ffmpeg -i input.mkv -c:v copy -c:a aac -strict experimental -b:a 320k -c:s mov_text -movflags +faststart output.mp4
# ffmpeg -i input.avi output.mp4
# find . -type f -name "*.mkv" -not -iwholename "*.AppleDouble*" -not -iwholename "*._*"
#
# Subtitles
# -c:s mov_text

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
