package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func ClientBase() {
	//address := "10.44.130.162:50000"
	address := "10.5.1.162:50000"
	for {

		start := time.Now()

		conn, err := net.Dial("tcp", address)

		end := time.Now().Sub(start)
		if err != nil {
			fmt.Println("Error dial:", err.Error())
			continue
		}

		now := time.Now()
		bytes := fmt.Sprint(now)
		_, err = conn.Write([]byte(bytes))

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(end, "ok")
		}

		var s int

		if len(os.Args) > 1 {
			fmt.Sscanf(os.Args[1], "%d", &s)
		} else {
			s = 1000
		}

		time.Sleep(time.Millisecond * time.Duration(s))
	}
}

func main() {
	fmt.Println("TCP client")
	ClientBase()
}
