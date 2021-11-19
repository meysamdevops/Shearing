package main

import (
	"fmt"
	"net"
	"shearing/Tcp_con"
	"shearing/cmd"
)

func main() {

	data := cmd.Tcp_cmd()

	fmt.Println(data)

	switch data.Check {
	case true:
		fmt.Println("Run as server")
		Server(data.Ip, data.Port)
	case false:
		fmt.Println("Run as client")
		Client(data.Ip, data.Port)
	}

}

func Server(ip string, port string) {
	server, err := net.Listen("tcp", ip+":"+port)
	if err != nil {
		panic(err)
	}

	defer server.Close()

	conn, err := server.Accept()
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	Tcp_con.Receive_file(conn, "my.mp", "3411968")
}

func Client(ip string, port string) {
	conn, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	Tcp_con.Send_file(conn, "test.mp4")
}
