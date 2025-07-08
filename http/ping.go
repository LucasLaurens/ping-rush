package web

import (
	"errors"
	"net/http"
	must "ping-rush/exception"
)

func Pings(urls []string) int {
	response := 200

	for _, url := range urls {
		response = must.Must(Ping(url))
	}

	return response
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
