package main

import "fmt"

func main()  {
	s:= "abcd"
	switch s[0] {
	case 'a':
		fmt.Println("The integer was <= 4")
		fallthrough
	case 'b':
		fmt.Println("The integer was <= 5")
		fallthrough
	case 'c':
		fmt.Println("The integer was <= 6")
		fallthrough
	default:
		fmt.Println("default case")
	}
}