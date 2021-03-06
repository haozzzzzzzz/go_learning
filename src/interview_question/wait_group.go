package main

import (
	"runtime"
	"sync"
	"fmt"
)

func main()  {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)

	for i := 0; i < 10; i ++ {
		go func() {
			fmt.Println("i:", i)
			wg.Done()
		}()
	}

	for i := 0; i < 10; i ++ {
		go func() {
			fmt.Println("i:", i)
			wg.Done()
		}()
	}

	wg.Wait()
}
/**
i: 10
i: 10
i: 10
i: 10
i: 10
i: 10
i: 10
i: 10
i: 10
i: 10
i: 10
i: 10
i: 10
i: 10
i: 10
i: 10
i: 10
i: 10
i: 10
i: 10
 */
