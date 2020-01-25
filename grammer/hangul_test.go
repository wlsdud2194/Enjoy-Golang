package main

import "fmt"

func ExampleHasConsonantSuffix() {
	fmt.Println(HasConsonantSuffix("Go 언어"))
	fmt.Println(HasConsonantSuffix("클로바"))
	fmt.Println(HasConsonantSuffix("미니언즈"))
	fmt.Println(HasConsonantSuffix("그럼"))
	// Output:
	// false
	// false
	// false
	// true
}
