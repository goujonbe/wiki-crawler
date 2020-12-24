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
		<h1>This is a Heading</h1>
		<p>This is a paragraph.</p>
		<a href="/wiki/test.com">my test</a>
		</div>
		<a href="https://fr.wikipedia.com/test">my second test</a>
		</body>
		</html>
	`)
	got := GetAllLinks(fakeReader)
	want := []string{"/wiki/test.com", "https://fr.wikipedia.com/test"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestRemoveDuplicates(t *testing.T) {
	testInput := []string{"a", "a", "b"}
	got := RemoveDuplicates(testInput)
	want := []string{"a", "b"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}
