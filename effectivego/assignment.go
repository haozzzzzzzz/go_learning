package main

import "fmt"

func echo() (int, int) {
	return 2, 3
}

func main() {
	v := 1
	a, v := echo()
	fmt.Println(v)
	fmt.Println(a)
}
