package main

import "fmt"

func fibonacci() func() int {
	idx := 0
	var vals []int
	var val int

	fn := func() int {
		val = 0
		if idx < 2 {
			val = 1
		} else {
			pre2 := vals[idx-2]
			pre1 := vals[idx-1]
			val = pre2 + pre1
		}

		vals = append(vals, val)
		idx++
		return val
	}

	return fn
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(i, f())
	}
}
