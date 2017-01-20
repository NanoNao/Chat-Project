package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
)

func main()  {
	conn, err := net.Dial("tcp", "127.0.0.1:666")
	if err != nil {

	}
	go handle(conn)
	for {

		Mess , _ := bufio.NewReader(os.Stdin).ReadString('\n')
		fmt.Fprintf(conn, Mess+"\n")

	}
}

func handle(conn net.Conn){
	for {
		Message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Printf(Message)
	}
}