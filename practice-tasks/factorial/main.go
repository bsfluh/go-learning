package main

import "fmt"

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	res := 1
	for i := 1; i <= x; i++ {
		res *= i
	}

	return res
}
func main() {
	fmt.Println(factorial(2))
}
