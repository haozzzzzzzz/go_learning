package main

import (
	"net/http"
	"unsafe"
)

const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a)
)

const (
	AX = 1
	BX = 2
	CX = 3
)

type A struct {
	XX   string
	YYAA string
}

func main() {

	a := &A{
		XX:   "2233",
		YYAA: "4455",
	}

	http.NewServeMux()
	println(a, b, c)
}
