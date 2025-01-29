// @Author Gopher
// @Date 2025/1/29 15:29:00
// @Desc
package main

import "fmt"

type Pet interface {
	eat()
	sleep()
}

type Dog struct {
	name string
	age  int
}

func (dog Dog) eat() {
	fmt.Println("dog eat...")
}

func (dog Dog) sleep() {
	fmt.Println("dog sleep...")
}

type Cat struct {
	name string
	age  int
}

func (cat Cat) eat() {
	fmt.Println("cat eat...")
}

func (cat Cat) sleep() {
	fmt.Println("cat amatelasi...")
}

type Person struct {
	name string
}

func (per Person) care(pet Pet) {

	pet.sleep()
	pet.eat()

}

func main() {
	dog := Dog{}
	cat := Cat{}
	per := Person{}

	per.care(dog)
	per.care(cat)
}
