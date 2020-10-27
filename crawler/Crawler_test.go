package crawler

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"parses/mocks"
	"strings"
	"testing"
)


func TestCrawler_Crawl(t *testing.T) {
	mockURL := "http://127.0.0.1"
	mockError := fmt.Errorf("mock error")
	mockReader := strings.NewReader("123")
	mockUrls := []string{"http://127.0.0.1/1", "http://127.0.0.1/2"}
	testData :=  map[string]struct{
		url string
		expected map[string]int
		parallelConnection int
		recDepth int
		err error
		fetcher *mocks.Fetcher
		parser *mocks.Parser
	}{
		"empty": {
			url: "",
			expected: nil,
		},
		"empty parallelConnection": {
			url: mockURL,
			expected: nil,
			err: mockError,
		},
		"less than 0 recDepth": {
			url: mockURL,
			expected: map[string]int{
				mockURL: 0,
			},
			recDepth: -1,
			parallelConnection: 1,
		},
		"recursionZero": {
			url: mockURL,
			expected: map[string]int{
				mockURL: 0,
			},
			recDepth: 0,
			parallelConnection: 10,
		},
		"Results return ": {
			url: mockURL,
			expected: map[string]int{
				mockURL: 0,
			},
			recDepth: 4,
			parallelConnection: 10,
			fetcher: func() *mocks.Fetcher {
				m := &mocks.Fetcher{}
				m.On("Fetch", mockURL).Once().Return(nil, mockError)
				return m
			}(),
		},
		"Fetcher's error is returned": {
			url: mockURL,
			expected: map[string]int{
				mockURL: 0,
			},
			recDepth: 2,
			parallelConnection: 10,
			fetcher: func() *mocks.Fetcher {
				m := &mocks.Fetcher{}
				m.On("Fetch", mockURL).Once().Return(nil, mockError)
				return m
			}(),
		},
		"Parser's error is returned": {
			url: mockURL,
			expected: map[string]int{
				mockURL: 0,
			},
			recDepth: 2,
			parallelConnection: 10,
			fetcher: func() *mocks.Fetcher {
				m := &mocks.Fetcher{}
				m.On("Fetch", mockURL).Once().Return(mockReader, nil)
				return m
			}(),
			parser: func() *mocks.Parser {
				p:= &mocks.Parser{}
				p.On("Parse", mockReader).Once().Return(nil, mockError)
				return p
			}(),
		},
		"All result": {
			url: mockURL,
			expected: map[string]int{
				mockURL: 0,
				"http://127.0.0.1/1":1,
				"http://127.0.0.1/2":1,
			},
			recDepth: 3,
			parallelConnection: 10,
			fetcher: func() *mocks.Fetcher {
				m := &mocks.Fetcher{}
				m.On("Fetch", mock.Anything).Times(3).Return(mockReader, nil)
				return m
			}(),
			parser: func() *mocks.Parser {
				p:= &mocks.Parser{}
				p.On("Parse", mockReader).Times(3).Return(mockUrls, nil)
				return p
			}(),
		},
	}
	for testName, tt := range testData {
		t.Run(testName, func(t *testing.T) {
			c := NewCrawler(tt.recDepth, tt.parallelConnection, tt.fetcher, tt.parser)
			result, err := c.Crawl(tt.url)

			if tt.err == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}

			assert.Equal(t, tt.expected, result)
			if tt.fetcher != nil {
				tt.fetcher.AssertExpectations(t)
			}
			if tt.parser != nil {
				tt.fetcher.AssertExpectations(t)
			}
		})
	}
}
