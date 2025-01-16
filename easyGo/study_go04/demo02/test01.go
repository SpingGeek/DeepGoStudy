package main

import "fmt"

type Profile struct {
	name string
	age  int
}

func NewProfile() *Profile {
	name := "张三"
	age := 20
	p := Profile{
		name: name,
		age:  age,
	}
	return &p
}

func main() {
	p := NewProfile()
	fmt.Println(p.name)
	fmt.Println(p.age)
}
