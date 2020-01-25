package main

import "fmt"

func main() {
	var (
		src   = []int{1, 2, 3, 4, 5}
		desc1 = make([]int, len(src))
		desc2 = append([]int(nil), src...)
	)

	if n := copy(desc1, src); n != len(src) {
		fmt.Println("not complete copy")
	} else {
		fmt.Println("complete copy")
	}
}
