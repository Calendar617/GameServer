package main

import (
    "fmt"
    "net"
    "os"
//    "net/http"
    //"io/ioutil"
)

//func say(w http.ResponseWriter, r *http.Request){
//	w.Write([]byte("pong"))
//	fmt.Println("say")
//	ioutil.WriteFile("say.txt", []byte("pong"), 0600)
//}

func main(){
	fmt.Println("Server Start")
	
	netListen, err := net.Listen("tcp", "localhost:8080")
	if err != nil { 
	  fmt.Println(os.Stderr, "Error : %s", err.Error())
	  return
	}
	
	defer netListen.Close()
	//服务器初始化成功
	fmt.Println("Server Init Success")
	//服务器开始监听
	fmt.Println("Server Listen Client")
	
	for{
		conn,err := netListen.Accept()
		if err != nil {
			continue
		}
		
	    fmt.Println("Client connect Server", conn.RemoteAddr());
		addclient(conn)
	}
	
	return
	
	//http.Handle("/", http.FileServer(http.Dir("./src")))
	//fmt.Println(http.Dir("./src"))
	//http.HandleFunc("/say", say)
	//if err :=http.ListenAndServe(":8080", nil); err != nil {
	//  panic(err)
	//}
}