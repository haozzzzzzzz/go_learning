package main

import "fmt"

func main() {
	type GameMapType [4][4]uint32

	var numbers [][]int
	numbers = make([][]int, 4)
	for i := 0; i < 4; i++ {
		numbers[i] = make([]int, 4)
	}

}
