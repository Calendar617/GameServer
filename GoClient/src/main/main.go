package main

import (
    "fmt"
    "net"
    "os"
)

func main(){
	fmt.Println("Client Start")
	connect()
}

func connect(){
	ipaddr := "127.0.0.1:8080"
    tcpAddr,err := net.ResolveTCPAddr("tcp4", ipaddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "clinet resolve addr faild Fatal error: %s", err.Error())
		return
	}
	
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "clinet connect faild Fatal error: %s", err.Error())
		return
	}
	
	fmt.Println("Client Connect Server Success", ipaddr)
	conn.Write([]byte("hello world"))
	go HandleConn(conn)
	Chat(conn)
}

func HandleConn(conn net.Conn){
	buffer := make([]byte,2048)
	for{
	   data, err := conn.Read(buffer)
	   if err != nil {
	   	fmt.Println("Client disconnect", conn.RemoteAddr());	
	   	 return
	   }
	   
	   fmt.Println(conn.RemoteAddr().String(), "receive data string:\n", string(buffer[:data])) 
	}
}

func Chat(conn net.Conn){
	fmt.Println("Client Can Chat")
	for{
		var input string 
		fmt.Scanln(&input)
		fmt.Println("Client scan message is", input)
		if conn != nil {
		  conn.Write([]byte(input))
		  fmt.Println("Client send message is", input)
		}
	}
}
