package pathflag

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

// Path is a custom flag var.
type Path struct {
	Path     string
	FileInfo os.FileInfo
}

// String returns the string value of Path.
func (path *Path) String() string {
	return path.Path
}

// Set the value of Path.
// Expand home directory if applicable.
// Make path absolute.
// Ensure file or directory is readable and writeable.
func (path *Path) Set(val string) error {
	if strings.HasPrefix(val, "~") {
		user, err := user.Current()
		if err != nil {
			return fmt.Errorf("Unable to expand home directory")
		}

		val = filepath.Join(user.HomeDir, val[2:])
	}

	val, err := filepath.Abs(val)
	if err != nil {
		return err
	}

	fileinfo, err := os.Stat(val)
	if err != nil {
		return err
	}

	// @todo: Get directory, check that is readable/writeable.

	perm := fmt.Sprintf("%s", fileinfo.Mode().Perm())[1:4]
	if !strings.Contains(perm, "r") || !strings.Contains(perm, "w") {
		return fmt.Errorf("Not readable or writeable")
	}

	path.Path = val
	path.FileInfo = fileinfo

	return nil
}
