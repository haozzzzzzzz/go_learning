package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var items = []int {
		1, 2, 3,
	}
	fmt.Println(items)
	rand.Seed(time.Now().UnixNano())
	randIndex := rand.Intn(len(items))
	fmt.Println(randIndex)
	fmt.Println(items[randIndex])
}
