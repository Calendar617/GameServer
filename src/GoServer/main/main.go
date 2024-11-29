package main

import (
	"fmt"
	"goserver/db"
	"net"
	"os"
	"time"
	//	"net/http"
	//
	// "io/ioutil"
)

// var wg sync.WaitGroup
var stopCh = make(chan bool, 1)

//func say(w http.ResponseWriter, r *http.Request){
//	w.Write([]byte("pong"))
//	fmt.Println("say")
//	ioutil.WriteFile("say.txt", []byte("pong"), 0600)
//}

var dbp *db.DBProcessor

func main() {
	fmt.Println("Server Start")
	//初始化Mysql
	dbp = db.CreateDBProcessor()
	//服务器开始监听
	netListen, err := net.Listen("tcp", "localhost:8880")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error : %s", err.Error())
		return
	}
	//服务器开始监听
	fmt.Println("Server Start Listen Client")
	defer netListen.Close()

	//服务器初始化成功
	fmt.Println("Server Init Success")
	//wg.Add(1)
	go Command()

	go func() {
		<-stopCh
		fmt.Println("server exit...")
		dbp.Close()
		time.Sleep(2 * time.Second)
		//wg.Wait()
		os.Exit(0)
	}()

	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}
		addClient(conn)
	}

	//http.Handle("/", http.FileServer(http.Dir("./src")))
	//fmt.Println(http.Dir("./src"))
	//http.HandleFunc("/say", say)
	//if err :=http.ListenAndServe(":8080", nil); err != nil {
	//  panic(err)
	//}
}

func Command() {
	//defer wg.Done()
	for {
		var input string
		fmt.Scanln(&input)
		//fmt.Println("Client scan message is", input)
		if input == "all" {
			fmt.Println("show all client info")
			showAllClient()
		} else if input == "exit" {
			fmt.Println("server start exit")
			stopCh <- true
			break
		}
	}
}
