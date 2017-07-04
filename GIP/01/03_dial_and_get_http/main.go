package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	/* net.Dial function connects using the type and endpoint specified
	   net.Dial returns connection and error, here we have ignored error by assigning it to _
	   godoc net Dial*/
	conn, _ := net.Dial("tcp", "golang.org:80")
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(status)
}
