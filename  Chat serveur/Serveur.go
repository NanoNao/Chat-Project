package main

import (
	"fmt"
	"net"
	"bufio"
	"strconv"
	"color"
)
type Color int
const (
	// No change of color

	Black=0
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)
var Test []net.Conn
func main()  {

	ct.Foreground(2, false)
	fmt.Println("Green text starts here...")

	ln, err := net.Listen("tcp", ":666")
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go handleConnection(conn)
	}

}
func handleConnection(conn net.Conn) {
	Test=append(Test,conn)
	var i int =0
	for i==0{
		Message, _ := bufio.NewReader(conn).ReadString('\n')

		fmt.Printf("-" + Message)
		fmt.Println(len(Test))
		for i:=0 ; i<len(Test) ;i++{
			for k:=0;k<len(Test) ;k++{
				if(Test[k]==conn){
					fmt.Fprintf(Test[i],"[Client "+strconv.Itoa(k)+"]- "+Message)
				}
			}

		}
		if Message == ""{
			i=1
			fmt.Printf("Disconnected \n")

		}

	}
}