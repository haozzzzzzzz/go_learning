package main

import "fmt"

func main() {
	var i int = 1

LOOP:
	for {
		switch {
		case i < 100:
			i++
			fmt.Println(i)

		case i > 100:
			break LOOP

		}
	}

}
