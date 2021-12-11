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
		Client(data.Ip, data.Port, data.File)
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

	massage := Tcp_con.Receive_massage(conn, 10)

	switch massage {
	case "file------":
		fmt.Println("Receive file")
		Tcp_con.S_First_con(conn)
	case "massage---":
		fmt.Println("Receive massage")
	default:
		fmt.Println("Receive Nothing")
	}

}

func Client(ip string, port string, Filename string) {
	conn, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	Tcp_con.Send_message(conn, "file------")
	Tcp_con.C_First_con(conn, Filename)
}
