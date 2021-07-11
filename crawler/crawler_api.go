package crawler

import "fmt"

func (crawl *crawler) Crawl(url string, depth int, c chan int) {

	if depth <= 0 {
		c <- depth
		return
	}

	if fetched := fetcher.checkFetchedUrls(url); fetched {
		fmt.Printf("URL: %s already fetched. \n", url)
		c <- -1
		return
	}

	body, urls, err := fetcher.fetch(url)
	fetcher.setFetchedUrls(url)
	if err != nil {
		fmt.Println(err)
		c <- -1
		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	innerc := make(chan int, len(urls))
	for _, u := range urls {
		go crawl.Crawl(u, depth-1, innerc)
	}

	for i := 0; i < len(urls); i++ {
		<-innerc
		// fmt.Println("Channel Data", depth, res)
	}

	// fmt.Println("reached crawl end", depth)
	c <- depth
}

func New() CrawlAPI {
	crawl := crawler{}
	return &crawl

}
