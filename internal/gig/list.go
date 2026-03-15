package gig

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

func (g *Gig) showList() error {
	resp, err := http.Get(gitignoreBaseURL)
	if err != nil {
		return errors.Wrapf(err, "failed to access URL: %s", gitignoreBaseURL)
	}
	defer func() { _ = resp.Body.Close() }()

	langCh := make(chan string)
	errCh := make(chan error, 1)
	go func() {
		errCh <- getLang(resp.Body, langCh)
		close(langCh)
	}()

	for v := range langCh {
		decoded, err := url.QueryUnescape(v)
		if err != nil {
			return errors.Wrapf(err, "failed to unescape: %s", v)
		}
		_, _ = fmt.Fprintln(g.OutStream, decoded)
	}

	if err := <-errCh; err != nil {
		return err
	}
	return nil
}

func getLang(r io.Reader, ch chan string) error {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return errors.Wrap(err, "failed to get document")
	}

	seen := make(map[string]bool)
	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if ok && strings.HasSuffix(href, gitignoreExt) && !seen[href] {
			seen[href] = true
			ch <- extractLang(href)
		}
	})
	return nil
}

func extractLang(s string) string {
	str := strings.Split(s, "/")
	return strings.ReplaceAll(str[len(str)-1], gitignoreExt, "")
}
