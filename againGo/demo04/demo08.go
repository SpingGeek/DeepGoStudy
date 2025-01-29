// @Author Gopher
// @Date 2025/1/29 16:13:00
// @Desc
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func responseSize(url string) {
	fmt.Println("Step1:", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Step2:", url)
	defer response.Body.Close()

	fmt.Println("Step3:", url)
	//fmt.Println("Response Size:", response.ContentLength)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Step4:", url)
	fmt.Println("Response Size:", len(body))

}
func main() {
	go responseSize("https://www.duoke360.com")
	go responseSize("https://baidu.com")
	go responseSize("https://jd.com")
	time.Sleep(10 * time.Second)
}
