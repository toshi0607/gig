package gig

import (
	"github.com/jessevdk/go-flags"
	"os"
	"fmt"
)

type config struct {
	List  bool `short:"l" long:"list" description:"shows list of available language"`
	File  bool `short:"f" long:"File" description:"outputs .ignore file"`
	Quiet bool `short:"q" long:"quiet" description:"hide stdout"`
	Args struct {
		Language string
	} `positional-args:"yes"`
}

func (g *Gig) initConfig() error {
	_, err := flags.ParseArgs(&g.Config, os.Args[1:])
	if err != nil {
		return err
	}

	if g.Config.List == false && g.Config.Args.Language == "" {
		fmt.Println("Usage: go run main.go <language> [options]")
		os.Exit(-1)
	}

	if g.Config.Quiet && !g.Config.File {
		fmt.Println("gig: output something!")
		os.Exit(-1)
	}
}
