package main

import "fmt"

type Vertex struct {
	x, y int
}

func main() {
	//var x uint64 = 1<<64 - 1
	//fmt.Printf("%d %x; %d %x \n", x, x, int64(x), int64(x))
	//
	//

	v := Vertex{1, 2}
	fmt.Printf("%v", v)

	m := map[string]int{
		"hello": 1,
	}
	fmt.Printf("%v", m)

}
