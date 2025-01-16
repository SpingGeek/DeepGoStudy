package main

import "fmt"

type Profile01 struct {
	name string
	age  int
}

func NewProfile01(name string, age int) *Profile01 {
	// initial
	p := Profile01{
		name: name,
		age:  age,
	}
	return &p
}

func main() {
	p := NewProfile01("张三", 20)
	fmt.Println(p.name)
	fmt.Println(p.age)
}
