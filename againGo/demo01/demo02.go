// @Author Gopher
// @Date 2025/1/28 21:24:00
// @Desc
package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello World!")
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Server is listening on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
