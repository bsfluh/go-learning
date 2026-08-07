package main

import "fmt"

func sum(sl []int) int {
	n := 0
	for i := range sl {
		n += sl[i]
	}
	return n
}
func main() {
	sl := []int{2, 3, 4, 6, 10, 22}
	fmt.Println(sum(sl))
}
