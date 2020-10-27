package fetchers

import (
	"io"
	"net/http"
)

type httpFetcher struct {
	userAgent string
	httpClient *http.Client
}


func NewHTTPFetcher(userAgent string, httpClient *http.Client) *httpFetcher {
	return &httpFetcher{
		userAgent:  userAgent,
		httpClient: httpClient,
	}
}

func (f *httpFetcher) Fetch(url string) (io.Reader, error) {

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", f.userAgent)

	resp, err := f.httpClient.Do(req)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp.Body, nil
}