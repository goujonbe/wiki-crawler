package crawler

import (
	"reflect"
	"strings"
	"testing"
)

func TestGetAllLinks(t *testing.T) {
	fakeReader := strings.NewReader(`
		<!DOCTYPE html>
		<html lang="en">

		<body>
		
		<img src="img_la.jpg" alt="LA" style="width:100%">
		
		<div>
		<h1> Heading </h1>
		<p>This is a paragraph.</p>
		<a href="/wiki/test.com">my test</a>
		<a href="/wiki/test.com">my test</a>
		<a href="/wiki/anothertest.com">my test</a>
		<a href="https://test.com">bad</a>
		</div>
		<a href="https://fr.wikipedia.com/test">my second test</a>
		</body>
		</html>
	`)
	got := GetWikipediaPageSummary(fakeReader)
	if got.Title != "Heading" {
		t.Errorf("got %q want %q", got.Title, "Heading")
	}
	want := []string{"/wiki/anothertest.com", "/wiki/test.com"}
	if !reflect.DeepEqual(got.Links, want) {
		t.Errorf("got %q want %q", got.Links, want)
	}
}
