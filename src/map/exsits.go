package main

import "fmt"

func main() {
	vals := make(map[string]int)
	vals["xx"] = 1
	vals["xxx"] = 2

	fmt.Println(vals)
	fmt.Println(vals["ll"])
}
