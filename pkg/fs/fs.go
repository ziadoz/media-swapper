package fs

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/ziadoz/media-swapper/pkg/pathflag"
)

func GetSwappableFiles(src pathflag.Path) ([]string, error) {
	mode := src.FileInfo.Mode()
	files := []string{}

	if mode.IsDir() {
		err := filepath.Walk(src.Path, func(path string, fileinfo os.FileInfo, err error) error {
			if IsSwappable(path) && !IsSwapped(path) {
				files = append(files, path)
			}

			return err
		})

		return files, err
	}

	if mode.IsRegular() && IsSwappable(src.Path) {
		return []string{src.Path}, nil
	}

	return files, nil
}

func SwapExt(path, ext string) string {
	return strings.TrimSuffix(path, filepath.Ext(path)) + "." + ext
}

func IsSwappableVideo(path string) bool {
	return filepath.Ext(path) == ".mkv"
}

func IsSwappableAudio(path string) bool {
	return filepath.Ext(path) == ".m4a"
}

func IsSwappable(path string) bool {
	return IsSwappableVideo(path) || IsSwappableAudio(path)
}

func IsSwapped(file string) bool {
	var output string
	if IsSwappableVideo(file) {
		output = SwapExt(file, "mp4")
	} else if IsSwappableAudio(file) {
		output = SwapExt(file, "mp3")
	}

	_, err := os.Stat(output)
	return err == nil
}

func LocateBinary() (string, error) {
	for _, bin := range []string{"avconv", "ffmpeg"} {
		path, err := exec.LookPath(bin)
		if err != nil {
			continue
		}

		return path, nil
	}

	return "", errors.New("Could not locate avconv or ffmpeg binary")
}
