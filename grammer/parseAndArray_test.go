package main

import (
	"fmt"
	"strconv"
)

func ExampleParseAndArray() {
	var (
		i   int64
		f   float64
		s   string
		b   bool
		err error
	)

	i, err = strconv.ParseInt("2048", 16, 64)
	f, err = strconv.ParseFloat("3.14", 64)
	s = strconv.Itoa(25)
	b, err = strconv.ParseBool("true")

	fmt.Println(i, f, s, b, err)

	animals := [4]string{
		"lion",
		"dog",
		"monkey",
		"dragon",
	}

	for _, animal := range animals {
		fmt.Println(animal)
	}

	// Output:
	// 8264 3.14 25 true <nil>
	// lion
	// dog
	// monkey
	// dragon
}
