package gig

import (
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
	p := flags.NewParser(&g.Config, flags.HelpFlag)
	_, err := p.Parse()
	if err != nil {
		return err
	}

	if g.Config.Version {
		return fmt.Errorf("gig version %s\n", version)
	}

	if !g.Config.List && g.Config.Args.Language == "" {
		p.WriteHelp(g.ErrStream)
		return fmt.Errorf("\n please check usage above")
	}

	return nil
}
