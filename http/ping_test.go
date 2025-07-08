package web

import (
	"testing"
)

func TestPingServeralUrl(t *testing.T) {
	value := struct {
		name     string
		urls     []string
		expected int
	}{
		name: "several urls test case",
		urls: []string{
			"https://www.google.com/",
			"https://fr.wikipedia.org/",
		},
		expected: 200,
	}

	t.Run(value.name, func(t *testing.T) {
		results := Pings(value.urls)

		for url, actual := range results {
			if actual != value.expected {
				t.Errorf(
					"the '%v' url status code expected %d but received %d",
					url,
					value.expected,
					actual,
				)
			}
		}
	})
}

func TestPingUrl(t *testing.T) {
	value := struct {
		name     string
		url      string
		expected int
	}{
		name:     "google test case",
		url:      "https://www.google.com/",
		expected: 200,
	}

	t.Run(value.name, func(t *testing.T) {
		acutal, err := Ping(value.url)

		if err != nil {
			t.Fatalf("failed to GET %s: %v", value.url, err)
		}

		if acutal != value.expected {
			t.Errorf(
				"the value returned status code %d; expected %d",
				acutal,
				value.expected,
			)
		}
	})
}
