package main

import (
	"fmt"

	"github.com/elvisferns/web-crawling/crawler"
)

func main() {
	crawler := crawler.New()

	c := make(chan int)
	go crawler.Crawl("https://golang.org/", 4, c)
	<-c
	fmt.Println("reached end")
}
