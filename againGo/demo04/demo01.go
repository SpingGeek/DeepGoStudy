// @Author Gopher
// @Date 2025/1/29 14:47:00
// @Desc
package main

import "fmt"

func sayHello(name string) {
	fmt.Printf("Hello, %s", name)
}

func f1(name string, f func(string)) {
	f(name)
}

func main() {
	f1("wangchen", sayHello)
}
