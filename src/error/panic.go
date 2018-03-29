package main

import (
	"fmt"
	"time"
)

func main()  {
	defer func() {
		fmt.Println("2")
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容，bug
		}
		fmt.Println("3")
	}()

	f()
}

func f() {
	for {
		fmt.Println("1")
		panic("bug")
		fmt.Println("4") // 不会运行
		time.Sleep( time.Second)
	}
}
