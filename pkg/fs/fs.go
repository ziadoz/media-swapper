package fs

import (
	"path/filepath"
	"strings"

	"github.com/ziadoz/media-swapper/pkg/pathflag"
)

func Find(path pathflag.Path, ext string) ([]string, error) {
	if path.FileInfo.IsDir() {
		return filepath.Glob(filepath.Join(path.Path, "*."+ext))
	}

	return []string{path.Path}, nil
}

func SwapExt(path, ext string) string {
	return strings.TrimSuffix(path, filepath.Ext(path)) + "." + ext
}
