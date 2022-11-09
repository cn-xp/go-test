package tcp

import (
	"fmt"
	"net"
	"os"
)

func TcpDialMain() {
	fmt.Println("---tcpDial---")
	conn, err := net.Dial("tcp", "93.184.216.34:80")
	checkConnection(conn, err)
	conn, err = net.Dial("udp", "93.184.216.34:80")
	checkConnection(conn, err)
	//conn, err = net.Dial("tcp", "[2400:da00::dbf:0:100]:80")
	//checkConnection(conn, err)
}

func checkConnection(conn net.Conn, err error) {
	if err != nil {
		fmt.Println("error connecting:", err.Error())
		os.Exit(1)
	}
	fmt.Printf("connection is made with %v\n", conn)
}
