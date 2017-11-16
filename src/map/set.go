package main

import "fmt"

func main() {
	set := make(map[uint32]bool)
	set[66] = true
	set[77] = true

	fmt.Println(set)

	for i, _ := range set {
		fmt.Println(i)
	}

}
