package main

import "fmt"

func main() {

	a := 7
	b := 11
	b += a
	fmt.Println(b)

	b -= a
	fmt.Println(b)

	b *= a
	fmt.Println(b)

	b /= a
	fmt.Println(b)

	b %= a
	fmt.Println(b)

	b &= a
	fmt.Println(b)

	b |= a
	fmt.Println(b)

	b ^= a
	fmt.Println(b)

	b <<= a
	fmt.Println(b)

	b >>= a
	fmt.Println(b)

}
