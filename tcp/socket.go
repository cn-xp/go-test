package tcp

import (
	"fmt"
	"io"
	"net"
)

func TcpSocketMain() {
	fmt.Println("---tcpSocket---")
	var (
		host   = "www.example.org"
		port   = "80"
		remote = host + ":" + port
		msg    = "GET / HTTP/1.1\r\n" + "Host:www.example.org\r\n" + "Connection: close\r\n" + "\r\n\r\n"
		data   = make([]uint8, 409600)
		read   = true
		count  = 0
	)
	conn, err := net.Dial("tcp", remote)
	io.WriteString(conn, msg)
	for read {
		count, err = conn.Read(data)
		if err != nil {
			fmt.Println("error read:", err.Error())
		}
		read = (err == nil)
		fmt.Printf(string(data[:count]))
	}
	conn.Close()
}
