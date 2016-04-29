package main

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"sync"

	"github.com/Sirupsen/logrus"
	"github.com/jessevdk/go-flags"
	"github.com/kr/fs"
)

func main() {
	// parsing arguments
	var opts struct {
		Chunks   int    `short:"C" long:"chunks" description:"Amount of chunks."`
		Affinity int    `long:"affinity" description:"If set, ptar will only archive the files of a specific affinity.  Else, it will create create all the archives."`
		Create   bool   `short:"c" description:"Create new archives containing the specified items."`
		File     string `short:"f" description:"Read the archive from or write the archive to the specified file.  The filename can be - for standard input or standard output."`
		// Verbose []bool `short:"v" description:"Show verbose debug information"`
		// Extract bool   `short:"x" description:"Extract to disk from the archives"`
		// List    bool   `short:"t" description:"List archive contents to stdout"`
	}

	opts.Affinity = -1

	args, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}

	// checking arguments
	if !opts.Create {
		logrus.Fatalf("ptar only works for creating (-c) archives for now.")
	}

	if opts.File == "" {
		logrus.Fatalf("You need to specify a file (-f myarchive.tar)")
	}

	if len(args) == 0 {
		logrus.Fatalf("You need to specify at least one file/directory to add in the archive")
	}

	if opts.Chunks == 0 {
		opts.Chunks = runtime.NumCPU()
	}
	logrus.Debugf("Using %d chunks", opts.Chunks)

	if opts.Affinity == -1 { // parent mode
		if opts.File == "-" {
			logrus.Fatalf("Using stdout as target is not compatible with parent mode.  You need to set the --affinity of the process.")
		}

		// prepare sub-commands
		commands := make([]*exec.Cmd, opts.Chunks)
		for i := 0; i < opts.Chunks; i++ {
			childArgs := append([]string{
				fmt.Sprintf("--affinity=%d", i),
				"-c",
				"-f", fmt.Sprintf("%s.%d", opts.File, i),
				"-C", fmt.Sprintf("%d", opts.Chunks),
			}, args...)
			logrus.Infof("child(%d): running %s %v", i, os.Args[0], childArgs)
			commands[i] = exec.Command(os.Args[0], childArgs...)
			commands[i].Stdout = os.Stdout
			commands[i].Stderr = os.Stderr
		}

		// call sub-commands in parallel
		var wg sync.WaitGroup
		for i := 0; i < opts.Chunks; i++ {
			wg.Add(1)
			go func(idx int) {
				command := commands[idx]
				command.Run()
				wg.Done()
			}(i)
		}

		// wait for all commands to finish
		wg.Wait()

	} else { // child mode
		var f io.Writer
		var err error

		// initialize archive
		if opts.File == "-" { // stdout mode
			f = os.Stdout
		} else { // file mode
			f, err = os.Create(opts.File)
			if err != nil {
				logrus.Fatalf("Failed to create %q: %v", opts.File, err)
			}
		}
		archive := tar.NewWriter(f)
		defer archive.Close()

		// iterate over files
		idx := -1
		for _, path := range args {
			walker := fs.Walk(path)

			for walker.Step() {
				idx++
				if idx%opts.Chunks != opts.Affinity {
					continue
				}
				if err := walker.Err(); err != nil {
					logrus.Warnf("fs error: %v", err)
					continue
				}
				if walker.Stat().IsDir() {
					continue
				}
				if walker.Stat().Mode()&os.ModeSocket != 0 {
					continue
				}

				srcFile, err := os.Open(walker.Path())
				if err != nil {
					logrus.Fatalf("failed to open %q: %v", walker.Path(), err)
				}
				defer srcFile.Close()

				header := &tar.Header{
					Name: walker.Path(),
					Size: walker.Stat().Size(),
				}
				logrus.Infof("+ %s", walker.Path())
				if err := archive.WriteHeader(header); err != nil {
					logrus.Fatalf("failed to write header of %q: %v", walker.Path(), err)
				}

				if _, err := io.Copy(archive, srcFile); err != nil {
					logrus.Fatalf("failed to write body of %q: %v", walker.Path(), err)
				}
			}
		}
	}
}
