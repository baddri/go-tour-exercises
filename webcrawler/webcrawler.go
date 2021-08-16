package webcrawler

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
// i'm not sure when to stop crawler so i set the limit
func Crawl(url string, depth int, fetcher Fetcher, c chan string, done *SafeMap, lim int) {
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)

	size := done.Add(url)

	if body != "" {
		fmt.Printf("found: %s %q\n", url, body)
	}

	// FIXME: it will break the sync if you set random delay here.
	// url have to be send right after adding it to safemap
	// otherwise it will break the sync and raise panic
	// `send on closed channel`

	// how to fix (idea):
	// 1. send url to channel inside Add method
	// 2. maybe check if channel is closed before send it
	// 		see -> https://stackoverflow.com/questions/16105325/how-to-check-a-channel-is-closed-or-not-without-reading-it

	if size < lim {
		c <- url
	} else {
		c <- url
		close(c)
		return
	}
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, u := range urls {
		// Dont fetch the same url twice ✔️
		// NOTE: for a little delay after Check() it will fetch same url twice or more thou �‍♂️� ,
		// but likely it will never happened
		if ok := done.Check(u); !ok {
			// Fetch url in paralel ✔️
			go Crawl(u, depth-1, fetcher, c, done, lim)
		}
	}
}

type SafeMap struct {
	mu sync.Mutex
	V  map[string]bool
}

// Check return true if key already in SafeMap
func (m *SafeMap) Check(k string) bool {
	m.mu.Lock()
	_, ok := m.V[k]
	defer m.mu.Unlock()
	return ok
}

// Add add key to SafeMap and get size of the map
func (m *SafeMap) Add(k string) int {
	m.mu.Lock()
	m.V[k] = true
	defer m.mu.Unlock()
	return len(m.V)
}

// FakeFetcher is Fetcher that returns canned results.
type FakeFetcher map[string]*FakeResult

type FakeResult struct {
	Body string
	Urls []string
}

func (f FakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.Body, res.Urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}
