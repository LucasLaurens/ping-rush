package web

import (
	"errors"
	"net/http"
	must "ping-rush/exception"
)

// todo: use go routine
func Pings(urls []string) map[string]int {
	results := make(map[string]int)

	for _, url := range urls {
		status := must.Must(Ping(url))
		results[url] = status
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
