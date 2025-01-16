package main

import (
	"fmt"
)

type profile struct {
	name     string
	age      int
	interest []string
}

func main() {

	p1 := profile{name: "LeoChen", age: 20, interest: []string{"coding", "basketball"}}
	fmt.Println(p1.name)
	fmt.Println(p1.age)
	fmt.Println(p1.interest)

	var p2 profile
	p2.name = "LeoChen"
	p2.age = 20
	p2.interest = []string{"coding"}
	fmt.Println(p2.name)
	fmt.Println(p2.age)
	fmt.Println(p2.interest)

	p3 := new(profile)
	p3.name = "spring"
	p3.age = 20
	p3.interest = []string{"hello", "world"}
	fmt.Println(p3.name)
	fmt.Println(p3.age)
	fmt.Println(p3.interest)

	p4 := &profile{}
	p4.name = "spring"
	p4.age = 20
	p4.interest = []string{"hello", "world"}

	/*
		p1 := profile{name: "LeoChen", age: 20, interest: []string{"coding", "basketball"}}
		var p2 profile
		p3 := new(profile)
	*/
}
