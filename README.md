# Media Swapper ![](https://github.com/ziadoz/media-swapper/workflows/GoReleaser/badge.svg)
A simple Bash script to swap MKV to MP4 and M4A to MP3.

## Requirements
You'll need to have either [ffmpeg](https://ffmpeg.org/) or [avconv](https://libav.org/avconv.html) installed on your machine for this script to work.

## Usage
You can convert all the media in the current directory by just running the script:
```
media-swapper /path/to/videos
```

You can also specify the binary if required:
```
media-swapper /path/to/videos --bin=/path/to/ffmpeg-or-avconv
```

If you need to know the version you are using, there's a flag for that:
```
media-swapper --version
```
