package gig

import (
	"net/http"
	"fmt"
	"strings"
	"net/url"
	"io"

	"github.com/PuerkitoBio/goquery"
)

func showList() {
	r, err := http.Get(gitignoreBaseURL)
	if err != nil {
		fmt.Println(err)
		return
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
