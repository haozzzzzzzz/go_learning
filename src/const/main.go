package main

import "fmt"

func main() {
	const (
		a = iota
		b
		c
		d
	)

	const (
		e = iota
		f
		g
	)

	const (
		h = iota
		i = iota
	)

	const (
		j = iota
		k = iota
	)

	fmt.Println(g)
	fmt.Println(h, i)
	fmt.Println(j, k)
}
