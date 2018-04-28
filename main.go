package main

import (
	"fmt"
	"os"

	"github.com/toshi0607/gig/gig"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintf(os.Stderr, "Error:\n%s", err)
			os.Exit(1)
		}
	}()
	cli := &gig.Gig{OutStream: os.Stdout, ErrStream: os.Stderr}
	os.Exit(cli.Run())
}
