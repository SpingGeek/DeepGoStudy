package main

import "fmt"

func main() {

	/*
		panic: runtime error: invalid memory address or nil pointer dereference
		[signal SIGSEGV: segmentation violation code=0x2 addr=0x0 pc=0x104754aa0]

		goroutine 1 [running]:
		main.main()
			/Users/spring/workspace/gopath/src/study_go03/demo08/test02.go:8 +0x20
	*/
	// null pointer
	var num *int32
	*num = 10
	fmt.Println(num)
}
