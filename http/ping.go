package web

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	must "ping-rush/exception"
	"sync"
)

type Result struct {
	url  string
	code int
	err  error
}

func Pings(urls []string) []Result {
	// Define context to manage error from goroutines
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Prepare sync to wait until all goroutines are complete before continuing processing
	var wg sync.WaitGroup
	// Buffer to prevent blocking
	channel := make(chan Result, len(urls))
	// Empty slice with capacity for URLs
	results := make([]Result, 0, len(urls))

	// Add as many tasks as there are URLs
	wg.Add(len(urls))
	for index, url := range urls {
		go func(occurence int, url string) {
			// It indicates when the goroutine is finished
			defer wg.Done()

			fmt.Printf(
				"The %d processing task is running...\n",
				occurence,
			)

			select {
			// We exit if the context is canceled, so errors occurs
			case <-ctx.Done():
				return
			default:
				code, err := Ping(url)
				if err != nil {
					// If an error occurs when pinging an URL, everything is canceled
					cancel()
				}

				channel <- Result{
					url:  url,
					code: code,
					err:  err,
				}
			}
		}(index+1, url)
	}

	// Close the channel after all goroutines have finished
	go func() {
		wg.Wait()
		close(channel)
	}()

	for result := range channel {
		results = append(results, result)
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
