// @Author Gopher
// @Date 2025/1/29 08:14:00
// @Desc
package main

import "fmt"

type Greeter interface {
	Greet(name string) string
}

type Person struct {
	greeting string
}

func (p Person) Greet(name string) string {
	return fmt.Sprintf(p.greeting, name)
}

func main() {
	var greeter Greeter
	greeter = Person{"Hello, %s!"}
	fmt.Println(greeter.Greet("Gopher"))

	message := greeter.Greet("Gopher")
	fmt.Println(message)
}
