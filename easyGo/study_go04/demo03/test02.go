package main

import (
	"fmt"
	"time"
)

func a() {
	time.Sleep(3 * time.Second)
	fmt.Println(" it is a ")
}

func b() {
	time.Sleep(2 * time.Second)
	fmt.Println(" it is b ")
}

func c() {
	time.Sleep(1 * time.Second)
	fmt.Println(" it is c ")
}

func main() {
	go a()
	go b()
	go c()
	time.Sleep(5 * time.Second)
}
