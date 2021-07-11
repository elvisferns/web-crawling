package crawler

import "fmt"

var fetcher = fakeFetcher{
	fetchedUrls: make(map[string]int),
	fetchUrls: map[string]*fakeResult{
		"https://golang.org/": {
			"The Go Programming Language",
			[]string{
				"https://golang.org/pkg/",
				"https://golang.org/cmd/",
			},
		},
		"https://golang.org/pkg/": {
			"Packages",
			[]string{
				"https://golang.org/",
				"https://golang.org/cmd/",
				"https://golang.org/pkg/fmt/",
				"https://golang.org/pkg/os/",
			},
		},
		"https://golang.org/pkg/fmt/": {
			"Package fmt",
			[]string{
				"https://golang.org/",
				"https://golang.org/pkg/",
			},
		},
		"https://golang.org/pkg/os/": {
			"Package os",
			[]string{
				"https://golang.org/",
				"https://golang.org/pkg/",
			},
		},
	},
}

func (f *fakeFetcher) fetch(url string) (string, []string, error) {
	if res, ok := f.fetchUrls[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

func (f *fakeFetcher) setFetchedUrls(url string) {
	f.mu.Lock()
	f.fetchedUrls[url] = 1
	f.mu.Unlock()
}

func (f *fakeFetcher) checkFetchedUrls(url string) bool {
	f.mu.Lock()

	defer f.mu.Unlock()
	_, ok := f.fetchedUrls[url]
	return ok
}
