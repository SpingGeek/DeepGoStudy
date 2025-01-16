package main

import "fmt"

func exchangeValue(i, j *int) {
	k := *i
	*i = *j
	*j = k
}

func main() {
	x, y := 7, 11
	fmt.Println(x, y)
	exchangeValue(&x, &y)
	fmt.Println(x, y)
}
