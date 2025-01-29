// @Author Gopher
// @Date 2025/1/28 21:44:00
// @Desc
package main

import (
	"flag"
	"fmt"
)

func main() {

	name := flag.String("name", "World", "a name to say hello to")
	flag.Parse()
	fmt.Printf("Hello, %s!\n", *name)

}
