package web

import "testing"

func TestPingUrl(t *testing.T) {
	values := []struct {
		name string
		url  string
		code int
	}{
		{
			name: "google test case",
			url:  "https://www.google.com/",
			code: 200,
		},
	}

	for _, value := range values {
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
}
