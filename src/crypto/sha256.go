package main

import (
	"crypto/sha256"
	"fmt"
)

func main()  {
	str := "appkey=5f03a35d00ee52a21327ab048186a2c4&random=7226249334&time=1457336869&mobile=13788888888"
	h:=sha256.New()
	h.Write([]byte(str))
	code :=  fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(code)
}
