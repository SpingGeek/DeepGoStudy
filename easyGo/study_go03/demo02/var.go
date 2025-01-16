package main

import "fmt"

// variable
var (
	age        int
	name       string
	getMarried bool
	days       int = 365
)

// init
func init() {
	println("ensure the goal, I can seek and achieve it, like golang")
}

// Drive /drive the car
func Drive() {
	consumeFuel := 38.6
	fmt.Println(consumeFuel)
}

// Hello the variable
func Hello() {
	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(getMarried)
	fmt.Println(days)
}

func GetYearDates() (int, int) {
	return 12, 24
}
