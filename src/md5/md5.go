package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func main() {
	s := md5.Sum([]byte("6666"))
	fmt.Println(s)
	str := fmt.Sprintf("%x", s)
	fmt.Println(str)
	n2 := str[0:2]
	fmt.Println(n2)
	fmt.Println([]byte(n2))
	i, _ := strconv.ParseInt(n2, 16, 0)
	fmt.Println(i)
}
