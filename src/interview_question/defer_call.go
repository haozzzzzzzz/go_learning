package main

import (
	"fmt"
	"time"
)

func main()  {
	defer_call()
}

func defer_call() {
	defer func() {
		fmt.Println(time.Now().UnixNano(), "打印前")
		time.Sleep(3*time.Second)

	}()

	defer func() {
		fmt.Println(time.Now().UnixNano(), "打印中")
		time.Sleep(3*time.Second)

	}()

	defer func() {
		fmt.Println(time.Now().UnixNano(), "打印后")
		time.Sleep(3*time.Second)

	}()

	panic(fmt.Sprintf("%d %s", time.Now().UnixNano(), "触发异常"))
}
/**
输出结构
1521532635579005000 打印后
1521532638579569500 打印中
1521532641579595100 打印前
panic: 1521532635578004600 触发异常
 */