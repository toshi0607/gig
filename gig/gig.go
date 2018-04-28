package gig

import (
	"io"
	"net/http"
	"fmt"
	"os"
	"time"
)

type Gig struct {
	OutStream, ErrStream io.Writer
	Output []io.Writer
}

func (g *Gig) Run() int {
	if config.List {
		showList()
		return 0
	}

	if config.File {
		var writer io.WriteCloser
		writer, err := os.Create(gitignoreExt + time.Now().Format("2006-01-02-15:04:05")) // for test
		if err != nil {
			fmt.Println(err)
			return 1
		}
		g.Output = append(g.Output, writer)
		defer writer.Close()
	}
	if !config.Quiet {
		g.Output = append(g.Output, os.Stdout)
	}

	lang := config.Args.Language
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
