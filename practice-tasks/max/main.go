package main

import (
	"errors"
	"fmt"
)

func max(nums []int) (int, error) {
	if len(nums) == 0 {
		return 0, errors.New("empty slice")
	}
	m := nums[0]
	for i := range nums {
		if m < nums[i] {
			m = nums[i]
		}
	}
	return m, nil
}
func main() {
	nums := []int{}
	n, err := max(nums)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n)
}
