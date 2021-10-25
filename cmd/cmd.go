package cmd

import (
	"flag"
)

type Data struct {
	Ip    string
	Port  string
	Check bool
	File  string
}

func Tcp_cmd() Data {
	ip := flag.String("ip", "localhost", "Ip addres for connection")
	port := flag.String("port", "5656", "Port for connection")
	check := flag.Bool("check", false, "run as server or clinet")
	client_file := flag.String("file", "non", "File for sending to server")

	flag.Parse()

	data := Data{
		Ip:    *ip,
		Port:  *port,
		Check: *check,
		File:  *client_file,
	}

	return data
}
