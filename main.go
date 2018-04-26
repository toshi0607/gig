package main

import (
	"fmt"
	"net/http"
	"io"
	"os"
	"flag"
)

var Config struct {
	List string
}

func init() {
	//flag.StringVar(&Config.List, "l", "", "show language list")
	flag.Parse()

	if len(flag.Args()) != 1 {
		fmt.Println("Usage: go run main.go <language> [options]")
		os.Exit(-1)
	}
}

func main() {
	lang := flag.Args()[0]
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
