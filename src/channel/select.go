package main

import (
	"fmt"
	"sync"
	"time"
)

//import "fmt"

func main() {
	c1 := make(chan int8)
	waiter := sync.WaitGroup{}
	waiter.Add(1)
	go func() {
		defer waiter.Done()

		for {
			select {
			case <-c1:
				fmt.Println("c1")
				//return
			default:
				fmt.Println("xxx")
				time.Sleep(time.Second)
			}
		}
	}()
	c1 <- 1
	waiter.Wait()
}
