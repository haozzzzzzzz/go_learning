package main

import "fmt"

func main() {
	arr := [2][2]int{
		{1,1},
		{2,2},
	}
	fmt.Println(arr)
	var arr2 [2][2]int = arr
	fmt.Println(arr2)

	arr2[1][1] = 3
	fmt.Println(arr2)
	fmt.Println(arr)
}
