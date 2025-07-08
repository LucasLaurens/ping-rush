package web

import "testing"

func TestPingServeralUrl(t *testing.T) {
	value := struct {
		name string
		urls []string
		code int
	}{
		name: "several urls test case",
		urls: []string{
			"https://www.google.com/",
			"https://fr.wikipedia.org/",
		},
		code: 200,
	}

	t.Run(value.name, func(t *testing.T) {
		code := Pings(value.urls)

		if code != value.code {
			t.Errorf(
				"the value returned status code %d; expected %d",
				code,
				value.code,
			)
		}
	})
}

func TestPingUrl(t *testing.T) {
	value := struct {
		name string
		url  string
		code int
	}{
		name: "google test case",
		url:  "https://www.google.com/",
		code: 200,
	}

	t.Run(value.name, func(t *testing.T) {
		code, err := Ping(value.url)

		if err != nil {
			t.Fatalf("failed to GET %s: %v", value.url, err)
		}

		if code != value.code {
			t.Errorf(
				"the value returned status code %d; expected %d",
				code,
				value.code,
			)
		}
	})
}

// todo: test bad url
