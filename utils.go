package main

import "sort"

type UrlVisited struct {
	url string
	depth int
}

type byDepth []*UrlVisited

func (r byDepth) Len() int           { return len(r) }
func (r byDepth) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r byDepth) Less(i, j int) bool { return r[i].depth < r[j].depth }


func UrlVisitedArrayFromMap(raw map[string]int) []*UrlVisited {
	results := make([]*UrlVisited, len(raw))
	i := 0
	for k, v := range raw {
		results[i] = &UrlVisited{
			url: k,
			depth: v,
		}
		i++
	}
	return results
}

func SortUrlVisitedByDepth(urlsVisited []*UrlVisited) []*UrlVisited {
	sort.Sort(byDepth(urlsVisited))
	return urlsVisited
}