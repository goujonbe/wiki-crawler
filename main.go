package main

import (
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	url := "https://en.wikipedia.org/wiki/Diego_Maradona"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	for _, link := range getAllLinks(resp.Body) {
		fmt.Println(link)
	}
}

func getAllLinks(r io.Reader) []string {
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
