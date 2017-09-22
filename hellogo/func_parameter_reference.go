package main

import "fmt"

func swap(x *int, y *int) {
	var temp int
	temp = *x //保留x值
	*x = *y   // 将x指针指向的值改为y指针指向的值
	*y = temp // 将temp赋值给y指针指向的值

}

func main() {
	var a int = 100
	var b int = 200

	fmt.Printf("交换前，a的值:%d\n", a)
	fmt.Printf("交换前，b的值:%d\n", b)

	swap(&a, &b)

	fmt.Printf("交换后，a的值:%d\n", a)
	fmt.Printf("交换后，b的值:%d\n", b)

}
