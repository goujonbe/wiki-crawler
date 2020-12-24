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
	allLinks := c.GetAllLinks(resp.Body)
	for _, link := range c.RemoveDuplicates(allLinks) {
		fmt.Println(link)
	}
}
