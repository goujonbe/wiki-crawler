package crawler

import (
	"io"
	"sort"
	"strings"

	"golang.org/x/net/html"
)

func GetAllLinks(r io.Reader) []string {
	var links []string
	z := html.NewTokenizer(r)
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return links
		case html.StartTagToken, html.EndTagToken:
			token := z.Token()
			if "a" == token.Data {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						links = append(links, attr.Val)
					}

				}
			}
		}
	}
}

// we are using the in place deduplicate algorithm
// as suggested in the official go wiki
// https://github.com/golang/go/wiki/SliceTricks#in-place-deduplicate-comparable
func RemoveDuplicates(s []string) []string {
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
func KeepWikiUrls(allUrls []string) []string {
	wikiUrls := []string{}
	for _, url := range allUrls {
		if strings.HasPrefix(url, "/wiki/") {
			wikiUrls = append(wikiUrls, url)
		}
	}
	return wikiUrls
}
