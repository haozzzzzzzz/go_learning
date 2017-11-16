package main

import (
	"time"
	"fmt"
)

func main()  {
	t1 := time.Now()
	time.Sleep(time.Second)
	t2 := time.Now()
	dur := t2.Sub(t1)
	fmt.Println(int64(float64(dur) / float64(time.Second)))
}
