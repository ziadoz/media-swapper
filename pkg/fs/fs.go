package fs

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/ziadoz/media-swapper/pkg/pathflag"
)

func GetSwappableFiles(src pathflag.Path) ([]string, error) {
	mode := src.FileInfo.Mode()
	files := []string{}

	if mode.IsDir() {
		err := filepath.Walk(src.Path, func(path string, fileinfo os.FileInfo, err error) error {
			if IsSwappable(path) {
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

func IsSwappableVideo(file string) bool {
	return filepath.Ext(file) == ".mkv"
}

func IsSwappableAudio(file string) bool {
	return filepath.Ext(file) == ".m4a"
}

func IsSwappable(file string) bool {
	return IsSwappableVideo(file) || IsSwappableAudio(file)
}
