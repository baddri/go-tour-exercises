package webcrawler

import (
	"fmt"
	"testing"
)

func TestWebcrawler(t *testing.T) {
	c := make(chan string)
	done := SafeMap{V: make(map[string]bool)}
	go Crawl("https://golang.org/", 4, fetcher, c, &done, 5)
	for {
		v, ok := <-c
		fmt.Println("channel c receives ", v)
		if !ok {
			return
		}
	}
}

var fetcher = FakeFetcher{
	"https://golang.org/": &FakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &FakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &FakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &FakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
