package crawler

import "sync"

// type fetcherAPI interface {
// 	fetch(url string) (body string, urls []string, err error)
// 	setFetchedUrls(url string)
// 	checkFetchedUrls(url string) bool
// }

type fakeFetcher struct {
	mu          sync.Mutex
	fetchUrls   map[string]*fakeResult
	fetchedUrls map[string]int
}

type fakeResult struct {
	body string
	urls []string
}
