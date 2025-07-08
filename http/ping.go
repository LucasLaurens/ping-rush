package web

import (
	"errors"
	"net/http"
	must "ping-rush/exception"
)

type Result struct {
	url  string
	code int
}

func Pings(urls []string) []Result {
	channel := make(chan Result, len(urls))

	for _, url := range urls {
		go func(url string) {
			channel <- Result{
				url:  url,
				code: must.Must(Ping(url)),
			}
		}(url)
	}

	results := make([]Result, 0, len(urls))
	for i := 0; i < len(urls); i++ {
		result := <-channel
		results = append(results, Result{
			url:  result.url,
			code: result.code,
		})
	}

	return results
}

func Ping(url string) (int, error) {
	if url == "" {
		return http.StatusBadRequest, errors.New(
			"the url cannot be empty",
		)
	}

	response := must.Must(http.Get(url))
	defer response.Body.Close()

	return response.StatusCode, nil
}
