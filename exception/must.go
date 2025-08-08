package must

import "fmt"

func Must[T any](value T, err error) T {
	if err != nil {
		fmt.Printf("unable to manage the value: %v", value)
		panic(err)
	}

	return value
}
