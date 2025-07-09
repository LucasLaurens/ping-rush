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
			"https://go.dev/doc/",
			"https://www.bbc.com/",
			"https://www.apple.com/",
			"https://www.twitch.tv/",
			"https://www.adobe.com/",
			"https://www.google.com/",
			"https://www.amazon.com/",
			"https://www.lemonde.fr/",
			"https://www.reddit.com/",
			"https://www.github.com/",
			"https://www.medium.com/",
			"https://www.dropbox.com/",
			"https://www.youtube.com/",
			"https://www.facebook.com/",
			"https://www.linkedin.com/",
			"https://www.wikipedia.org/",
			"https://www.microsoft.com/",
			"https://www.instagram.com/",
			"https://www.salesforce.com/",
			"https://www.stackoverflow.com/",
		},
		expected: 200,
	}

	t.Run(value.name, func(t *testing.T) {
		results := Pings(value.urls)

		for _, result := range results {
			if result.code != value.expected {
				t.Errorf(
					"the %v url expected %d status code but received %d",
					result.url,
					value.expected,
					result.code,
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
