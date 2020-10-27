package main

import (
	"flag"
	"fmt"
	"net/http"
	"parses/crawler"
	"parses/fetchers"
	"parses/parsers"
)

func main () {
	root := flag.String("root", "https://meduza.io/", "начальная страница")
	n := flag.Int("n", 100, "максимальное количество параллельных запросов")
	r := flag.Int("r", 3, "глубина рекурсии")
	userAgent := flag.String("user-agent",
		"Mozilla/5.0 (Android 8.0.0; Mobile; rv:61.0) Gecko/61.0 Firefox/68.0", "заголовок User-Agent")

	flag.Parse()

	f := fetchers.NewHTTPFetcher(*userAgent, http.DefaultClient)
	p := parsers.NewHTMLParser()

	c := crawler.NewCrawler(*r, *n, f, p)
	res, _ := c.Crawl(*root)

	results := SortUrlVisitedByDepth(UrlVisitedArrayFromMap(res))

	fmt.Println(len(results))

	for _, urlVisited := range results {
		fmt.Println(urlVisited.depth, urlVisited.url)
	}

}
