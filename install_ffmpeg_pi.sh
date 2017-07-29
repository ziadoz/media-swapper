#!/usr/bin/env bash

# Install FFMpeg on Raspberry Pi
wget https://johnvansickle.com/ffmpeg/releases/ffmpeg-release-armhf-32bit-static.tar.xz
tar xf ./ffmpeg-release-armhf-32bit-static.tar.xz
mv ./ffmpeg-3.3.2-armhf-32bit-static/ffmpeg /usr/bin
rm -rf ./ffmpeg-3.3.2-armhf-32bit-static/ ffmpeg-release-armhf-32bit-static.tar.xz