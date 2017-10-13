package main

import (
	"fmt"
	"net"
)

func ServerBase() {
	//address := "10.44.130.162:50000"
	address := "10.5.1.162:50000"
	fmt.Println("Starting the server...")
	//create listener
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}

	// listen and accept connections from clients:
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err.Error())
			return
		}
		//create a goroutine for each request.
		go doServerStuff(conn)
	}

}

func doServerStuff(conn net.Conn) {
	fmt.Println("new connection:", conn.LocalAddr())
	buf := make([]byte, 1024)
	length, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}

	fmt.Println("Receive data from client:", string(buf[:length]))
	conn.Close()
}

func main() {
	fmt.Println("TCP server")
	ServerBase()
}
