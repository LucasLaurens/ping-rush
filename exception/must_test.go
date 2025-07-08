package must

import (
	"testing"
)

func TestPingUrl(t *testing.T) {
	values := []struct {
		name   string
		actual any
		expect any
	}{
		{
			name:   "2 equal words",
			actual: "hello world",
			expect: "hello world",
		},
		{
			name:   "2 equal numbers",
			actual: 1,
			expect: 1,
		},
	}

	for _, value := range values {
		t.Run(value.name, func(t *testing.T) {
			if value.actual != value.expect {
				t.Errorf(
					"The actual value does not correspond to the excpeted one: %v, %v",
					value.actual,
					value.expect,
				)
			}
		})
	}
}
