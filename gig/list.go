package gig

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func (g *Gig) showList() error {
	r, err := http.Get(gitignoreBaseURL)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	langCh := make(chan string)
	go func() {
		getLang(r.Body, langCh)
		close(langCh)
	}()

	for v := range langCh {
		decoded, err := url.QueryUnescape(v)
		if err != nil {
			return err
		}
		fmt.Fprintln(g.OutStream, decoded)
	}
	return nil
}

func getLang(r io.Reader, ch chan string) error {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return err
	}

	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		url, ok := s.Attr("href")
		if ok && strings.HasSuffix(url, gitignoreExt) {
			ch <- extractLang(url)
		}
	})
	return nil
}

func extractLang(s string) string {
	str := strings.Split(s, "/")
	return strings.Replace(str[len(str)-1], gitignoreExt, "", -1)
}
