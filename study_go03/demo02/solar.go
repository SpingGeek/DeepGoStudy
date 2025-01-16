package main

import "fmt"

func getYearDates() (int, int) {
	return 12, 24
}
func main() {

	monthNums, _ := getYearDates()
	_, solarNums := getYearDates()
	fmt.Println(monthNums, solarNums)
	//fmt.Println(GetYearDates())
}
