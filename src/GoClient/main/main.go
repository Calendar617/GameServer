package main

import (
	"fmt"
	"goclient/logic"
	"net"
	"os"
	"strconv"
	"time"
)

var clientMap map[uint64]*Client = make(map[uint64]*Client)
var commandList = [2]string{"delete", "deleteall"}

// func startServer(addr chan string) {
// 	l, err := net.Listen("tcp", ":0")
// 	if err != nil {
// 		log.Fatal("network error", err)
// 	}
// 	addr <- l.Addr().String()
// }

func main() {
	// addr := make(chan string)
	// go startServer(addr)

	// conn, _ := net.Dial("tcp", <-addr)
	// defer func() {
	// 	_ = conn.Close()
	// }()

	fmt.Println("Client Start")

	var index = uint64(1)
	var input string
	fmt.Scanln(&input)
	index, error := strconv.ParseUint(input, 0, 64)
	//for index < 2 {
	if error == nil {
		connect(index)
	}
	//	index++
	//}

	for {
		time.Sleep(1000)
	}
}

func connect(index uint64) {
	ipaddr := "127.0.0.1:8880"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ipaddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "client resolve addr faild Fatal error: %s", err.Error())
		return
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "client connect faild Fatal error: %s", err.Error())
		return
	}

	time.Sleep(1000)
	client := CreateClient(index, "", conn)
	clientMap[index] = client
	go client.Chat()
	fmt.Println("Client Connect Server Success, My id is ", index)

	logic.Testlogic(100)
}

func isCommand(cmd string) bool {
	for i := 0; i < len(commandList); i++ {
		if commandList[i] == cmd {
			return true
		}
	}
	return false
}

func Command(cmd []string) {
	fmt.Println("Client cmd is", cmd[0])
	if cmd[0] == "delete" {
		id, err := strconv.Atoi(cmd[1])
		fmt.Println("Client delete, id is ", id)
		if err == nil {
			c, exist := clientMap[uint64(id)]
			if exist {
				c.CloseConn()
				delete(clientMap, 0)
			} else {
				fmt.Println("Client delete, id not exist")
			}
		} else {
			fmt.Println("Client strconv.Atoi, error is ", err)
		}
	} else if cmd[0] == "deleteall" {

	}
}
