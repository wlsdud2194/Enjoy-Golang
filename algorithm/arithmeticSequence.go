/** 등차수열 */
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 첫째항 값, 등차, n차 항
	const a1, diff, n int = 1, 10, 5

	fmt.Printf("\n첫째항(a1): %d, 등차(d): %d, n차 항까지(n): %d\n", a1, diff, n)

	// ArithmeticalList(a1, diff, n)
	Sum := GetArithmeticalSum(a1, diff, n)
	fmt.Print(Sum)
}

/** 등차 수열의 합 */
func GetArithmeticalSum(a1, diff, section int) int {
	var result int = 0
	const method = 2

	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	n := random.Intn(method)
	fmt.Println("Method : ", n)

	if n == 0 {
		for i := 1; i <= section; i++ {
			result += (a1 + (i-1)*diff) // 등차 수열의 일반항 공식
		}
	} else if n == 1 {
		result = (section * (2*a1 + (section-1)*diff)) / 2 // 등차 수열의 합 공식
	}

	return result
}
