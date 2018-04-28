package gig

import (
	"io"
	"net/http"
	"fmt"
	"os"
	"time"
)

const version = "v0.1.0"

type Gig struct {
	OutStream, ErrStream io.Writer
	Output []io.Writer
	Config config
}

func (g *Gig) Run() int {
	err := g.initConfig()
	if err != nil {
		fmt.Println(err)
		return 1
	}

	if g.Config.List {
		err := showList()
		if err != nil {
			fmt.Println(err)
			return 1
		}
		return 0
	}

	if g.Config.File {
		var writer io.WriteCloser
		writer, err := os.Create(gitignoreExt + time.Now().Format("2006-01-02-15:04:05")) // for test
		if err != nil {
			fmt.Println(err)
			return 1
		}
		g.Output = append(g.Output, writer)
		defer writer.Close()
	}
	if !g.Config.Quiet {
		g.Output = append(g.Output, os.Stdout)
	}

	lang := g.Config.Args.Language
	url := gitignoreFileBaseURL + lang + gitignoreExt
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return 1
	}
	defer resp.Body.Close()

	dest := io.MultiWriter(g.Output...)
	_, err = io.Copy(dest, resp.Body)
	if err != nil {
		fmt.Println(err)
		return 1
	}
	return 0
}
