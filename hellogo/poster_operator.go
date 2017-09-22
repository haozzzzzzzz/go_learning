package main

import "fmt"

func main() {
	var a int = 4
	var b int32
	var c float32
	var ptr *int

	/*
		运算符实例
	*/
	fmt.Printf("第1行 - a变量类型为 = %T\n", a)
	fmt.Printf("第2行 - b变量类型为 = %T\n", b)
	fmt.Printf("第3行 - c变量类型为 = %T\n", c)
	fmt.Printf("第4行 - ptr变量类型为 = %T\n", ptr)

	/**
	& 和 * 运算符实例
	*/
	ptr = &a // ptr 包含了 a 变量的地址
	fmt.Printf("a的值为 %d\n", a)
	fmt.Printf("*ptr为%d\n", *ptr)

}
