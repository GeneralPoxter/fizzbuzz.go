package main

import (
	"strconv"
)

// Struct example
type cond struct {
	key int
	str string
}

// Variadic function example
func fizzbuzz(k, n int, conds ...cond) (out []string) {
	// Initialize string slice of size n
	out = make([]string, n)

	// Construct for loop
	for i := 0; i < n; i++ {
		// Declare variables
		var str string
		x := k + i

		// Iterate through range
		for _, cond := range conds {
			// If statement example
			if x%cond.key == 0 {
				str = str + cond.str
			}
		}

		if len(str) == 0 {
			// Int to string conversion
			str = strconv.Itoa(x)
		}

		out[i] = str
	}
	return out
}
