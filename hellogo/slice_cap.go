package main

import "fmt"

func main() {
	var numbers []int
	numbers = append(numbers, 0)
	printSlice(numbers)

	numbers = append(numbers, 1)
	printSlice(numbers)

	numbers = append(numbers, 2, 3, 4, 5)
	printSlice(numbers)

	numbers = append(numbers, 7, 8, 9)
	printSlice(numbers)

	// 创建切片numbers2是之前切片的两倍容量
	numbers2 := make([]int, len(numbers), (cap(numbers) * 2))
	copy(numbers2, numbers)
	printSlice(numbers2)

}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
