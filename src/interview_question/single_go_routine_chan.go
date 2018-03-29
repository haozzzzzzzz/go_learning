package main

import (
	"runtime"
	"fmt"
)

func main()  {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "Hello"
	select {
	case value := <-int_chan:
		panic(value)

	case value := <- string_chan:
		fmt.Println(value)
	}
}
/**
有时输出Hello，有时候panic
 */