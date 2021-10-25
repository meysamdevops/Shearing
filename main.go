package main

import (
	"fmt"
	"net"
	"os"
	"shearing/Tcp_con"
	"shearing/cmd"
)

func main() {

	data := cmd.Tcp_cmd()

	fmt.Println(data)

	switch data.Check {
	case false:
		err := Tcp_con.S_accessibility(data.Ip, data.Port)
		if err == nil {
			fmt.Println("Run as server...")
			fmt.Println("---------- server listening on " + data.Ip + " -----------")
			for {
				Server(data.Ip, data.Port)
			}
		} else {
			fmt.Println(err)
			os.Exit(1)
		}
	case true:
		fmt.Println("In client")
		err := Tcp_con.C_accessibility(data.Ip, data.Port)
		if err == nil {
			fmt.Println("Run as client...")
			Client(data.Ip, data.Port, data.File)
		} else {
			fmt.Println(err)
			os.Exit(1)
		}
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
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	fmt.Println("---------- Connected -----------")

	message := Tcp_con.Receive_massage(conn, 7)

	fmt.Println("Message:" + message)

	switch message {
	case "Contest":
		Tcp_con.Send_message(conn, "ConOk")
	case "Data---":
		Tcp_con.S_First_con(conn)
		Tcp_con.Send_message(conn, "Ok")
	}

}

func Client(ip string, port string, File string) {
	conn, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	Tcp_con.Send_message(conn, "Data---")
	Tcp_con.C_First_con(conn, File)

	fmt.Println(Tcp_con.Receive_massage(conn, 2))
}
