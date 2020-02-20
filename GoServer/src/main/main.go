package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    //"strings"
)

func say(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("pong"))
	fmt.Println("say")
	ioutil.WriteFile("say.txt", []byte("pong"), 0600)
}

func main(){
	http.Handle("/", http.FileServer(http.Dir("./src")))
	fmt.Println(http.Dir("./src"))
	http.HandleFunc("/say", say)
	if err :=http.ListenAndServe(":8080", nil); err != nil {
	  panic(err)
	}
}

