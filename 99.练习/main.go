// fib.go
package main

import (
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	rand.Intn(10)
	strings.ToLower("aa")
}

func mins(p ...int) int {
	min := p[0]
	for i := 0; i < len(p); i++ {
		if min > p[i] {
			min = p[i]
		}
	}
	return min
}
