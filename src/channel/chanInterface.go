package main

import "fmt"

func main() {

	chanInt := make(chan int)
	go func() {
		chanInt<-1
	}()
	newVal := <-chanInt
	fmt.Println(newVal)

}
