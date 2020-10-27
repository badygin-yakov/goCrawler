package parsers

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/url"
	"strings"
)

type HTMLParser struct {}

func NewHTMLParser() *HTMLParser {
	return &HTMLParser{}
}

func concatUrl(baseHref string, href string) string {
	base, _ := url.Parse(baseHref)
	hrefUrl, _ := url.Parse(href)

	return base.ResolveReference(hrefUrl).String()
}

func (p *HTMLParser) Parse(baseUrl string, reader io.Reader) ([]string, error){
	var res []string

	document, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}

	document.Find("a").Each(func(i int, element *goquery.Selection) {
		href, exists := element.Attr("href")
		if exists {
			if !strings.Contains(href, "https://") {
				href = concatUrl(baseUrl, href)
			}

			res = append(res, href)
		}
	})

	return res, err
}