package crawler

import (
	"fmt"
	"io"
	"log"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (reader io.Reader, err error)
}

type Parser interface {
	Parse(rootUrl string, reader io.Reader) (urls []string, err error)
}

type HistoryResult struct {
	url string
	recDepth int
}

type Crawler struct {
	recDepth           int
	parallelConnection int
	fetcher            Fetcher
	parser             Parser
}

func NewCrawler(recDepth int, parallelConnection int, fetcher Fetcher, parser Parser) *Crawler {
	return &Crawler{
		recDepth:           recDepth,
		parallelConnection: parallelConnection,
		fetcher:            fetcher,
		parser:             parser,
	}
}

func (c *Crawler) Crawl(url string) (map[string]int, error) {
	if url == "" {
		return nil, nil
	}

	if c.parallelConnection <= 0 {
		return nil, fmt.Errorf("need parallelConnection in init")
	}
	wg := sync.WaitGroup{}
	visitUrl := map[string]int {
		url: 0,
	}
	parsingResults := make(chan HistoryResult)
	urlTasks := make(chan HistoryResult)
	defer func() {
		close(parsingResults)
		close(urlTasks)
	}()
	go func() {
		for hr := range parsingResults {
			if _, ok := visitUrl[hr.url]; !ok {
				visitUrl[hr.url] = hr.recDepth
				wg.Add(1)
				urlTasks <- hr
			}
		}
	}()

	for i := 0; i < c.parallelConnection; i++ {
		go func() {
			for hr := range urlTasks {
				results, err := c.crawl(hr.url, hr.recDepth + 1)
				if err != nil {
					log.Println("Error: ", err)
				}
				for _, result := range results {
					parsingResults <- result
				}
				wg.Done()
			}
		}()
	}
	wg.Add(1)
	urlTasks <- HistoryResult{url, 0}
	wg.Wait()
	return visitUrl, nil
}

func (c *Crawler) crawl(url string,  depth int) ([]HistoryResult, error){

	if depth >= c.recDepth {
		return nil, nil
	}

	resp, err := c.fetcher.Fetch(url)
	if err != nil {
		return nil, nil
	}

	urls, err := c.parser.Parse(url, resp)
	if err != nil {
		return nil, nil
	}

	var results []HistoryResult
	for _, url := range urls {
		results = append(results, HistoryResult{url, depth})
	}

	return results, nil
}

