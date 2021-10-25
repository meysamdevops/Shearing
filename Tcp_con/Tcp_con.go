package Tcp_con

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

func Tk() {
	fmt.Println("Khorji")
}

func S_First_con(conn net.Conn) {
	data := bufio.NewReader(conn)

	re, _ := data.ReadString('\n')

	sp := strings.Split(re, ":")

	fmt.Println(sp)

	Receive_file(conn, sp[0], sp[1])
}

func C_First_con(conn net.Conn, filename string) {
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	info, err := file.Stat()
	if err != nil {
		panic(err)
	}

	fmt.Println(info.Name() + ":" + strconv.FormatInt(info.Size(), 10))
	Send_message(conn, info.Name()+":"+strconv.FormatInt(info.Size(), 10)+":\n")
	Send_file(conn, filename)

}

func Send_file(conn net.Conn, filename string) {
	file, _ := os.Open(filename)
	fi, _ := file.Stat()

	b := make([]byte, fi.Size())

	for {
		_, err := file.Read(b)
		if err == io.EOF {
			break
		}
		conn.Write(b)
	}
}

func Send_message(conn net.Conn, message string) {
	conn.Write([]byte(message))
}

func Receive_massage(conn net.Conn, count int) string {

	b := make([]byte, count)

	for {
		n, err := conn.Read(b)

		if n >= count {
			return string(b)
		}

		if err != nil {
			if err == io.EOF {
				return string(b)
			}
			return string(b)
		}
	}

}

func Receive_file(conn net.Conn, filename string, filesize string) {
	defer fmt.Println("Wait for File -----------------------------------------")
	fmt.Println("Start of receive " + filename)

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	defer file.Close()

	size, _ := strconv.Atoi(filesize)

	fmt.Println("file info: ", file.Name(), "   ", size)

	b := make([]byte, size)

	for {

		n, err := conn.Read(b)

		file.Write(b)
		//fmt.Println(b)
		if n >= size {
			break
		}
		if err != nil {
			if err == io.EOF {
				break
			}
		}

	}

	fmt.Println("Receive " + filename + " was completed")

}

func S_accessibility(ip string, port string) error {
	Listen, err := net.Listen("tcp", ip+":"+port)
	if err != nil {
		return err
	}
	defer Listen.Close()

	return nil
}

func C_accessibility(ip string, port string) error {
	fmt.Println("IN ACC")
	fmt.Println("Dial to " + ip + ":" + port)
	Listen, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		return err
	}
	defer Listen.Close()

	Send_message(Listen, "Contest")

	message := Receive_massage(Listen, 5)

	switch message {
	case "ConOk":
		return nil
	default:
		return fmt.Errorf("Have problem with connecting to server")
	}
}
