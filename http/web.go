package web

import (
	"errors"
	"net/http"
)

func Ping(url string) (int, error) {
	if url == "" {
		return 500, errors.New("empty URL")
	}

	response, err := http.Get(url)
	if err != nil {
		return response.StatusCode, err
	}

	defer response.Body.Close()

	return response.StatusCode, nil
}

// todo: Pings - several urls
