package main

import (
	"strings"
	"fmt"
)

func main() {
	index := strings.Index("http://baidu.com", "http")
	fmt.Println(index)
}
