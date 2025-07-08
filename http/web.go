package web

import (
	"errors"
	"net/http"
	must "ping-rush/exception"
)

// todo: Pings - several urls

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
