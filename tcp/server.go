package tcp

import (
	"fmt"
	"net"
	"os"
	"strings"
)

var mapUsers map[string]int

func TcpServerInit() {
	var listener net.Listener
	var err error
	var conn net.Conn
	mapUsers = make(map[string]int)

	fmt.Println("starting the server...")

	listener, err = net.Listen("tcp", "localhost:5000")
	checkServerError(err)
	for {
		conn, err = listener.Accept()
		checkServerError(err)
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	var buf []byte
	var err error
	for {
		buf = make([]byte, 512)
		_, err = conn.Read(buf)
		checkServerError(err)
		input := string(buf)
		if strings.Contains(input, "SH") {
			fmt.Println("server shuting down.")
			os.Exit(0)
		}
		if strings.Contains(input, "WHO") {
			displayList()
		}
		index := strings.Index(input, "says")
		clientName := input[0 : index-1]
		mapUsers[clientName] = 1
		fmt.Printf("received data: %v\n", string(buf))
	}
}

func displayList() {
	fmt.Println("----------------------")
	fmt.Println("this is the client list: 1=active, 0=inactive")
	for key, value := range mapUsers {
		fmt.Printf("user %s is %d\n", key, value)
	}
	fmt.Println("----------------------")
}

func checkServerError(err error) {
	if err != nil {
		panic("error:" + err.Error())
	}
}
