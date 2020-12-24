package crawler

import (
	"io"
	"sort"

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
