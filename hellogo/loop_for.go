package main

import "fmt"

func main() {
	var b int = 15
	var a int

	numbers := [6]int{1, 2, 3, 5}

	// for 循环
	for a := 0; a < 10; a++ {
		fmt.Printf("a的值为：%d\n", a)
	}

	for a < b {
		a++
		fmt.Printf("a的值为:%d\n", a)
	}

	for i, x := range numbers {
		fmt.Printf("第%d位x的值为=%d\n", i, x)
	}

}
