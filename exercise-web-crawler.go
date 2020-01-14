package main

import (
	"fmt"
	"sync"
)

// Fetcher ...
type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// FetchedUrls ...
type FetchedUrls struct {
	mux  sync.Mutex
	urls []string
}

// Set ...
func (c *FetchedUrls) Set(url string) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.urls = append(c.urls, url)
}

// IsCached ...
func (c *FetchedUrls) IsCached(url string) (bool, error) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for _, u := range c.urls {
		if u == url {
			return true, fmt.Errorf("already fetched. url: %s", url)
		}
	}
	return false, nil
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel. --> DONE
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:

	var wg sync.WaitGroup

	cache := FetchedUrls{}

	var crawler func(string, int, Fetcher)
	crawler = func(url string, depth int, fetcher Fetcher) {
		defer wg.Done()
		if depth <= 0 {
			return
		}

		_, err := cache.IsCached(url)
		if err != nil {
			fmt.Println("\t", err)
			return
		}
		cache.Set(url)
		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println("\t", err)
			return
		}

		fmt.Printf("found: %s %q\n", url, body)

		for _, u := range urls {
			wg.Add(1)
			go crawler(u, depth-1, fetcher)
		}
	}
	wg.Add(1)
	go crawler(url, depth, fetcher)
	wg.Wait()
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
