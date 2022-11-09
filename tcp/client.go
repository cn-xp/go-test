package tcp

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func TcpClientInit() {
	conn, err := net.Dial("tcp", "localhost:5000")
	checkError(err)
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("first,what is your name?")
	clientName, _ := inputReader.ReadString('\n')
	trimmedClient := strings.Trim(clientName, "\r\n")
	for {
		fmt.Println("what to send to the server? Type Q to quit.Type SH to shutdown server.")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")
		if trimmedInput == "Q" {
			return
		}
		_, err = conn.Write([]byte(trimmedClient + " says:" + trimmedInput))
		checkError(err)
	}
}

func checkError(err error) {
	if err != nil {
		panic("error:" + err.Error())
	}
}
