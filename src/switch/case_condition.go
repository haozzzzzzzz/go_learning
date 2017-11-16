package main

import "fmt"

func main() {
	var a = 5
	switch {
	case a >1 && a < 10:
		fmt.Println("1")
		break
	case a == 0:
		fmt.Println("2")
		break
	}
}
