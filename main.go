package main

import (
	"fmt"
	"net/http"

	c "wiki-crawler/crawler"
)

func main() {
	url := "https://en.wikipedia.org/wiki/Diego_Maradona"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	currentPage := c.GetWikipediaPageSummary(resp.Body)
	fmt.Println(currentPage.Title)
	fmt.Println(currentPage.Links[:5])
}
