package main

import (
     "fmt"
     "net"
     "strconv"
)

var clientMap map[int]net.Conn
var clientId = 0
func addclient(conn net.Conn){
	clientMap := make(map[int]net.Conn)
	clientMap[clientId] = conn
	go HandleConn(conn, clientId)
	clientId++
}

func HandleConn(conn net.Conn, id int){
	buffer := make([]byte,2048)
	for{
	   data, err := conn.Read(buffer)
	   if err != nil {
	   	fmt.Println("Client disconnect", conn.RemoteAddr());	
	   	 return
	   }
	   
	   fmt.Println(conn.RemoteAddr().String(), "receive data string:\n", string(buffer[:data]))
	   var context = "Welcom,your id is "
	   context += strconv.Itoa(id)
	   conn.Write([]byte(context))
	}
}
