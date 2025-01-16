package main

import (
	"log"
	"net/http"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("新年快乐 老爸老妈！"))
}

func main() {
	http.HandleFunc("/say_hello", SayHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Print("ListenAndServe: ", err)
		return
	}
}
