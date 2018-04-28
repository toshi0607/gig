package gig

import (
	"io"
	"net/http"
	"fmt"
	"os"
	"time"
	"strings"
	"net/url"

	"github.com/jessevdk/go-flags"
	"github.com/PuerkitoBio/goquery"

)

type Gig struct {
	OutStream, ErrStream io.Writer
}

var config struct {
	List  bool `short:"l" long:"list" description:"shows list of available language"`
	File  bool `short:"f" long:"File" description:"outputs .ignore file"`
	Quiet bool `short:"q" long:"quiet" description:"hide stdout"`
	Args struct {
		Language string
	} `positional-args:"yes"`
}

func init() {
	_, err := flags.ParseArgs(&config, os.Args[1:])
	if err != nil {
		fmt.Println(err)
		return
	}

	if config.List == false && config.Args.Language == "" {
		fmt.Println("Usage: go run main.go <language> [options]")
		os.Exit(-1)
	}

	if config.Quiet && !config.File {
		fmt.Println("gig: output something!")
		os.Exit(-1)
	}
}

func (g *Gig) Run() int {
	if config.List {
		r, err := http.Get(gitignoreBaseURL)
		if err != nil {
			fmt.Println(err)
			return 1
		}
		defer r.Body.Close()
		langCh := make(chan string)
		go func() {
			getLang(r.Body, langCh)
			close(langCh)
		}()

		for v := range langCh {
			fmt.Println(v)
		}

		return 1
	}

	var writers []io.Writer

	if config.File {
		var writer io.WriteCloser
		writer, err := os.Create(gitignoreExt + time.Now().Format("2006-01-02-15:04:05")) // for test
		if err != nil {
			fmt.Println(err)
			return 1
		}
		writers = append(writers, writer)
		defer writer.Close()
	}
	if !config.Quiet {
		writers = append(writers, os.Stdout)
	}

	lang := config.Args.Language
	url := gitignoreFileBaseURL + lang + gitignoreExt
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return 1
	}
	defer resp.Body.Close()

	writers = append(writers, os.Stdout)
	dest := io.MultiWriter(writers...)
	_, err = io.Copy(dest, resp.Body)
	if err != nil {
		fmt.Println(err)
		return 1
	}
	return 0
}

func getLang(r io.Reader, ch chan string) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		fmt.Println(err)
	}

	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		url, ok := s.Attr("href")
		if ok && strings.HasSuffix(url, gitignoreExt) {
			decoded, err := extractLang(url)
			if err != nil {
				fmt.Println(err)
			}
			ch <- decoded
		}
	})
}

func extractLang(s string) (string, error) {
	str := strings.Split(s, "/")
	return url.QueryUnescape(strings.Replace(str[len(str)-1], gitignoreExt, "", -1))
}
