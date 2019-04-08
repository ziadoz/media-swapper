# Media Swapper [![Build Status](https://travis-ci.org/ziadoz/media-swapper.svg?branch=master)](https://travis-ci.org/ziadoz/media-swapper)
A simple Bash script to swap MKV to MP4 and M4A to MP3.

## Requirements
You'll need to have either [ffmpeg](https://ffmpeg.org/) or [avconv](https://libav.org/avconv.html) installed on your machine for this script to work.

## Usage
You can convert all the media in the current directory by just running the script:
```
./media-swapper --src=/path/to/videos --bin=/path/to/ffmpeg-or-avconv
```

You can also specify the binary if required:
```
./media-swapper --src=/path/to/videos --bin=/path/to/ffmpeg-or-avconv
```
