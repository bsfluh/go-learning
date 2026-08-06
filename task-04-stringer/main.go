package main

import "fmt"

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}
func main() {
	hosts := map[string]IPAddr{
		"first_test":  {2, 4, 55, 7},
		"second_test": {44, 55, 234, 4},
	}
	for name, ip := range hosts {
		fmt.Printf("%v : %v\n", name, ip)
	}

}
