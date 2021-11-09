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
func Crawl(url string, depth int, fetcher Fetcher, data *SafeMap, wg *sync.WaitGroup) {
	defer wg.Done()

	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	data.Add(url)

	if body != "" {
		fmt.Printf("found: %s %q\n", url, body)
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, u := range urls {
		// Dont fetch the same url twice ✔️
		if ok := data.Check(u); !ok {
			// Fetch url in paralel ✔️
			wg.Add(1)
			go Crawl(u, depth-1, fetcher, data, wg)
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
	defer m.mu.Unlock()
	_, ok := m.V[k]
	return ok
}

// Add add key to SafeMap and get size of the map
func (m *SafeMap) Add(k string) int {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.V[k] = true
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
