# web-crawling  

go application for concurrency with sync.Mutex

It implements a fake web crawling.

Using concurrency fake fetch requests are sent to different url  
Also mutex is used to avoid mutliple calls to same url

## Prerequisites

* install go

## Steps to build

**go build** will build and create an executable web-crawling in the root folder  
It will also download the dependencies for the first time
