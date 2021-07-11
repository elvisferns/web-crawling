package crawler

type CrawlAPI interface {
	Crawl(url string, depth int, c chan int)
}

type crawler struct{}
