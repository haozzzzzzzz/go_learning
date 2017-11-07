package main

import (
	"fmt"
	"reflect"
)

func main() {
	//var signal chan int = make(chan int)
	//go func() {
	//	signal <- 1
	//}()
	//val := <-signal
	//fmt.Println(val)
	//

	var signal = make(chan interface{})
	var number = uint32(666)
	go func() {
		signal <- 1
		signal <- 0.2
		signal <- &number
	}()
	val := <-signal
	val2 := <-signal
	val3 := <-signal
	fmt.Println(reflect.TypeOf(val), val)
	fmt.Println(reflect.TypeOf(val2), val2)
	fmt.Println(reflect.TypeOf(val3), *val3.(*uint32))

	go func() {
		signal <- 6
	}()
	val4 := <- signal
	fmt.Println(val4)
}


