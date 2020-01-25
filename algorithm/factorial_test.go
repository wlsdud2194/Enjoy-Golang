package main

import (
	"fmt"
)

func factorial(num int) int {
	if num <= 0 {
		return 1
	}

	return num * factorial(num-1)
}

func ExampleFactorial() {
	fmt.Println(factorial(4))
	fmt.Println(factorial(5))
}
