package main

import "fmt"

func sum(a []int, c chan int) {
	fmt.Println(a)
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // 从c中获取
	fmt.Println(x, y, x+y)
}
