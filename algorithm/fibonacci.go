package main

import (
	"fmt"
)

func main() {
	p, q := 0, 1
	const n int = 10

	for i := 0; i <= n; i++ {
		fmt.Print(p)
		fmt.Print(" ")
		p, q = q, p+q
	}
}
