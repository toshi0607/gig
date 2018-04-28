package gig

import (
	"os"
	"fmt"

	"github.com/jessevdk/go-flags"
)

type config struct {
	List  bool `short:"l" long:"list" description:"shows list of available language"`
	File  bool `short:"f" long:"File" description:"outputs .ignore file"`
	Quiet bool `short:"q" long:"quiet" description:"hides stdout"`
	Version bool `short:"v" long:"version" description:"shows version"`
	Args struct {
		Language string
	} `positional-args:"yes"`
}

func (g *Gig) initConfig() error {
	_, err := flags.ParseArgs(&g.Config, os.Args[1:])
	if err != nil {
		return err
	}

	if g.Config.Version {
		fmt.Printf("gig version %s\n", version)
		os.Exit(-1)
	}

	if g.Config.List == false && g.Config.Args.Language == "" {
		fmt.Println("Usage: go run main.go <language> [options]")
		os.Exit(-1)
	}

	if g.Config.Quiet && !g.Config.File {
		fmt.Println("gig: output something!")
		os.Exit(-1)
	}

	return nil
}
