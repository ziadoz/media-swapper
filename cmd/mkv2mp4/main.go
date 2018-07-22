// Swap MKV container files to MP4 easily.
// Usage: mkv2mp4 --src=/path/to/videos --bin=/path/to/ffmpeg
package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/ziadoz/media-swapper/pkg/fs"
	"github.com/ziadoz/media-swapper/pkg/mp4swap"
	"github.com/ziadoz/media-swapper/pkg/pathflag"
)

var bin pathflag.Path
var src pathflag.Path

type result struct {
	cmd *mp4swap.Cmd
	err error
}

func init() {
	flag.Var(&bin, "bin", "The location of the ffmpeg or avconv binary")
	flag.Var(&src, "src", "The source directory of mkvs or individual mkv file to swap to mp4")
	flag.Parse()
}

func main() {
	if bin.Path == "" {
		fmt.Fprintln(os.Stderr, "The -bin flag must be specified")
		os.Exit(1)
	}

	if src.Path == "" {
		fmt.Fprintln(os.Stderr, "The -src flag must be specified")
		os.Exit(1)
	}

	files, err := fs.Find(src, "mkv")
	if err != nil || len(files) == 0 {
		fmt.Fprintf(os.Stderr, "Could not find mkv files: %s\n", err)
		os.Exit(1)
	}

	in := make(chan *mp4swap.Cmd)
	out := make(chan *result)

	go queueCmds(files, in)
	go processCmds(in, out)

	fmt.Printf("Swapping %d videos: \n", len(files))
	for result := range out {
		if result.err != nil {
			fmt.Printf(" - Failed: %s: %s\n", result.cmd.Input, result.err)
		} else {
			fmt.Printf(" - Swapped: %s\n", result.cmd.Input)
		}
	}
}

func queueCmds(files []string, in chan *mp4swap.Cmd) {
	for _, input := range files {
		in <- mp4swap.Command(bin.Path, input)
	}

	close(in)
}

func processCmds(in chan *mp4swap.Cmd, out chan *result) {
	for cmd := range in {
		var cmdout, cmderr bytes.Buffer
		cmd.Stdout = &cmdout
		cmd.Stderr = &cmderr

		var reserr error
		if err := cmd.Run(); err != nil {
			reserr = err
			if strings.Contains(cmderr.String(), "already exists. Overwrite ? [y/N]") {
				reserr = fmt.Errorf("mp4 file already exists")
			}
		}

		out <- &result{
			cmd: cmd,
			err: reserr,
		}
	}

	close(out)
}
