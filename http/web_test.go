package web

import "testing"

type Website struct {
	url  string
	code int
}

var websites []Website

func TestPingUrl(t *testing.T) {
	websites = []Website{
		{
			url:  "https://www.google.com/",
			code: 200,
		},
		// {
		// 	url:  "",
		// 	code: 500,
		// },
	}

	for _, website := range websites {
		StatusCode, Err := Ping(website.url)

		if Err != nil {
			t.Fatalf("failed to GET %s: %v", website.url, Err)
		}

		// defer response.Body.Close()

		if StatusCode != website.code {
			t.Errorf(
				"the website returned status code %d; expected %d",
				StatusCode,
				website.code,
			)
		}
	}
}
