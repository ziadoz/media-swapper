package swap

import (
	"os/exec"

	"github.com/ziadoz/media-swapper/pkg/fs"
)

type Cmd struct {
	*exec.Cmd
	Input  string
	Output string
}

// Command returns a Cmd to convert an MKV to an MP4.
func Mp4Command(binary string, input string) *Cmd {
	output := fs.SwapExt(input, "mp4")
	args := []string{
		"-i",
		input,
		"-nostats",
		"-loglevel",
		"0",
		"-c:v",
		"copy",
		"-c:a",
		"copy",
		"-c:s",
		"mov_text",
		"-movflags",
		"+faststart",
		//"-nostdin",
		output,
	}

	return &Cmd{
		Cmd:    exec.Command(binary, args...),
		Input:  input,
		Output: output,
	}
}

// Command returns a Cmd to convert an M4A to an MP3.
func Mp3Command(binary string, input string) *Cmd {
	output := fs.SwapExt(input, "mp3")
	args := []string{
		"-i",
		input,
		"-nostats",
		"-loglevel",
		"0",
		"-c:v",
		"copy",
		"-c:a",
		"libmp3lame",
		"-q:a",
		"2",
		//"-nostdin",
		output,
	}

	return &Cmd{
		Cmd:    exec.Command(binary, args...),
		Input:  input,
		Output: output,
	}
}
