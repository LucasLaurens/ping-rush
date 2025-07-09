package web

import (
	"errors"
	"fmt"
	"net/http"
	must "ping-rush/exception"
)

type Result struct {
	url  string
	code int
}

func Pings(urls []string) []Result {
	channel := make(chan Result)
	results := make([]Result, len(urls))

	for index, url := range urls {
		// todo: manage error from goroutine to stop the process during execution
		// note: use sync.Group
		go func(occurence int, url string) {
			fmt.Printf(
				"The %d processing task is running...\n",
				occurence,
			)

			channel <- Result{
				url:  url,
				code: must.Must(Ping(url)),
			}
		}(index+1, url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-channel
		results[i] = Result{
			url:  result.url,
			code: result.code,
		}
	}

	fmt.Println("The processing task is complete !")
	return results
}

// note: should i pass with pointer on receiver ?
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
