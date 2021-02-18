package main

import (
	"strconv"
)

type cond struct {
	key int
	str string
}

func fizzbuzz(k, n int, conds ...cond) (out []string) {
	out = make([]string, n)
	for i := 0; i < n; i++ {
		var str string
		x := k + i
		for _, cond := range conds {
			if x%cond.key == 0 {
				str = str + cond.str
			}
		}

		if len(str) == 0 {
			str = strconv.Itoa(x)
		}

		out[i] = str
	}
	return out
}
