package main

import "fmt"

func main() {

	rankList := [4]string{"chen", "leo", "spring", "golang"}
	for i, v := range rankList {
		fmt.Printf("第 %d 名 %v\n", i+1, v)
		// in the language of golang , using the array is not frequent
	}
}
