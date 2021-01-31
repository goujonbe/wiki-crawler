package crawler

import (
	"io"
	"sort"
	"strings"

	"golang.org/x/net/html"
)

type WikipediaPage struct {
	Title string
	Links []string
}

// Parses HTML content. Extracts title and unique Wikipedia links
func GetWikipediaPageSummary(r io.Reader) WikipediaPage {
	page := WikipediaPage{}
	z := html.NewTokenizer(r)
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			page.Links = removeDuplicates(page.Links)
			page.Links = keepWikiUrls(page.Links)
			return page
		case html.StartTagToken, html.EndTagToken:
			token := z.Token()
			if "a" == token.Data {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						page.Links = append(page.Links, attr.Val)
					}
				}
			}
			if "h1" == token.Data {
				if tt = z.Next(); tt == html.TextToken {
					data := strings.TrimSpace(z.Token().Data)
					if len(data) > 0 {
						page.Title = data
					}
				}
			}
		}
	}
}

// we are using the in place deduplicate algorithm
// as suggested in the official go wiki
// https://github.com/golang/go/wiki/SliceTricks#in-place-deduplicate-comparable
func removeDuplicates(s []string) []string {
	sort.Strings(s)
	j := 0
	for i := 1; i < len(s); i++ {
		if s[j] == s[i] {
			continue
		}
		j++
		s[j] = s[i]
	}
	return s[:j+1]
}

// keep URLs that start with /wiki/ and remove the others
// that can link to external websites (in particular in the ref section)
func keepWikiUrls(allUrls []string) []string {
	wikiUrls := []string{}
	for _, url := range allUrls {
		if strings.HasPrefix(url, "/wiki/") {
			wikiUrls = append(wikiUrls, url)
		}
	}
	return wikiUrls
}
