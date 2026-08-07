package main

import "fmt"

func fibonacci(n int) []int {
	a := 0
	b := 1
	sl := make([]int, n)
	for i := 0; i < n; i++ {

		sl[i] = a
		a, b = b, a+b //изначально написал a,b=b,a+1 и результат был ошибочным, думал отм чтобы сдела a+b, но подумал, что это неверно и подсмотрел как решал до этого
	}
	return sl
}
func main() {
	fmt.Println(fibonacci(10))
}
