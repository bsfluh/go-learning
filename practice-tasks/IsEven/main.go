package main

import "fmt"

func isEven(x int) bool {
	return x%2 == 0
}
func main() {
	fmt.Println(isEven(4))
}
