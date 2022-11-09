package tcp

import (
	"flag"
	"fmt"
	"net"
	"syscall"
)

const maxRead = 25

func SimpleTcpServerMain() {
	flag.Parse()
	if flag.NArg() != 2 {
		panic("usage: host port")
	}
	hostAndPort := fmt.Sprintf("%s:%s", flag.Arg(0), flag.Arg(1))
	listener := initServer(hostAndPort)
	for {
		conn, err := listener.Accept()
		checkTcpError(err, "Accept:")
		go connectionHandle(conn)
	}
}

func initServer(hostAndPort string) *net.TCPListener {
	serverAddr, err := net.ResolveTCPAddr("tcp", hostAndPort)
	checkTcpError(err, "resolving address:port failed: '"+hostAndPort+"'")
	listener, err := net.ListenTCP("tcp", serverAddr)
	checkTcpError(err, "ListenTCP:")
	fmt.Println("Listening to: ", listener.Addr().String())
	return listener
}

func connectionHandle(conn net.Conn) {
	connFrom := conn.RemoteAddr().String()
	fmt.Println("connection from: ", connFrom)
	sayHello(conn)
	for {
		var ibuf []byte = make([]byte, maxRead+1)
		length, err := conn.Read(ibuf[:maxRead])
		ibuf[maxRead] = 0
		switch err {
		case nil:
			handleMsg(length, err, ibuf)
		case syscall.EAGAIN:
			continue
		default:
			goto DISCONNECT
		}
	}
DISCONNECT:
	err := conn.Close()
	fmt.Println("closing connection: ", connFrom)
	checkTcpError(err, "close:")
}

func sayHello(to net.Conn) {
	obuf := []byte{'L', 'e', 't', '\'', 's', ' ', 'G', '0', '!', '\n'}
	wrote, err := to.Write(obuf)
	checkTcpError(err, "write: wrote "+string(wrote)+" bytes.")
}

func handleMsg(length int, err error, msg []byte) {
	if length > 0 {
		fmt.Print("<", length, ":")
		for i := 0; ; i++ {
			if msg[i] == 0 {
				break
			}
			fmt.Printf("%c", msg[i])
		}
		fmt.Print(">")
	}
}

func checkTcpError(err error, info string) {
	if err != nil {
		panic("error:" + info + " " + err.Error())
	}
}
