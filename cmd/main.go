// Swap MKV/M4A containers to MP4/MP3.
// Usage: media-swapper --src=/path/to/videos --bin=/path/to/ffmpeg-or-avconv
package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/ziadoz/media-swapper/pkg/fs"
	"github.com/ziadoz/media-swapper/pkg/pathflag"
	"github.com/ziadoz/media-swapper/pkg/swap"
)

const workers int = 4

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

var (
	ver  bool
	help bool
	bin  pathflag.Path
)

var usage = `Media Swapper is a simple tool for container swapping m4a audio to mp3, and mkv video to mp4.

Usage: media-swapper [--bin=<path>] <path>
       media-swapper --help
       media-swapper --version`

type result struct {
	cmd *swap.Cmd
	err error
}

func main() {
	flag.BoolVar(&ver, "version", false, "The version of media swapper")
	flag.BoolVar(&help, "help", false, "Show command usage instructions and help")
	flag.Var(&bin, "bin", "The location of the ffmpeg or avconv binary")
	flag.Parse()

	if help {
		fmt.Println(usage)
		os.Exit(0)
	}

	if ver {
		fmt.Printf("Version: %s\nCommit: %s\nDate: %s\n", version, commit, date)
		os.Exit(0)
	}

	if bin.Path == "" {
		path, _ := fs.LocateBinary()
		stat, _ := os.Stat(path)

		bin = pathflag.Path{
			Path:     path,
			FileInfo: stat,
		}
	}

	if bin.Path == "" {
		fmt.Fprintln(os.Stderr, "The -bin flag must be specified")
		os.Exit(1)
	}

	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "A source must be specified")
		os.Exit(1)
	}

	src := pathflag.Path{}
	src.Set(os.Args[1])

	files, err := fs.GetSwappableFiles(src)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not find mkv/m4a files: %s\n", err)
		os.Exit(1)
	}

	if len(files) == 0 {
		fmt.Println("No mkv/m4a files were found to swap")
		os.Exit(0)
	}

	fmt.Printf("Swapping %d files: \n", len(files))

	in := make(chan *swap.Cmd)
	out := make(chan *result)
	done := make(chan struct{})
	wg := sync.WaitGroup{}

	go results(out, done)
	go pool(&wg, workers, in, out)
	go queue(files, in)

	<-done
}

func queue(files []string, in chan *swap.Cmd) {
	for _, file := range files {
		if fs.IsSwappableVideo(file) {
			in <- swap.Mp4Command(bin.Path, file)
		} else if fs.IsSwappableAudio(file) {
			in <- swap.Mp3Command(bin.Path, file)
		}
	}

	close(in)
}

func pool(wg *sync.WaitGroup, workers int, in chan *swap.Cmd, out chan *result) {
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go worker(wg, in, out)
	}

	wg.Wait()
	close(out)
}

func worker(wg *sync.WaitGroup, in chan *swap.Cmd, out chan *result) {
	for cmd := range in {
		err := cmd.Run()

		out <- &result{
			cmd: cmd,
			err: err,
		}
	}

	wg.Done()
}

func results(out chan *result, done chan struct{}) {
	for result := range out {
		if result.err != nil {
			fmt.Printf(" - Failed: %s: %s\n", result.cmd.Input, result.err)
		} else {
			fmt.Printf(" - Swapped: %s\n", result.cmd.Input)
		}
	}

	done <- struct{}{}
}
