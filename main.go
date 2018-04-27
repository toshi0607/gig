package main

import (
	"fmt"
	"net/http"
	"io"
	"os"
	"strings"
	"net/url"

	"golang.org/x/net/html"

	"github.com/jessevdk/go-flags"
)

var config struct {
	List  bool  `short:"l" long:"list" description:"shows list of available language"`
	Args    struct {
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
}

func main() {
	if config.List {
		url := "https://github.com/github/gitignore"
		r, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer r.Body.Close()
		langCh := make(chan string)
		go func() {
			parseItem(r.Body, langCh)
			close(langCh)
		}()

		for v := range langCh {
			fmt.Println(v)
		}

		return
	}

	lang := config.Args.Language
	url := "https://raw.githubusercontent.com/github/gitignore/master/" + lang + ".gitignore"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func parseItem(r io.Reader, ch chan string) {
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Println(err)
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" && strings.HasSuffix(a.Val, ".gitignore") {
					s := strings.Split(a.Val, "/")
					decoded, err := url.QueryUnescape(strings.Replace(s[len(s)-1], ".gitignore", "", -1))
					if err != nil {
						fmt.Println(err)
					}
					ch <- decoded
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}
