// @Author Gopher
// @Date 2025/1/21 17:31:00
// @Desc
package main

import "fmt"

type LinkNode struct {
	Data     int64
	NextNode *LinkNode
}

func main() {
	// new node
	node := new(LinkNode)
	node.Data = 2

	// new node
	node1 := new(LinkNode)
	node1.Data = 3
	node.NextNode = node1

	// new node
	node2 := new(LinkNode)
	node2.Data = 4
	node1.NextNode = node2

	// sort the node to print
	nowNode := node
	for {
		if nowNode != nil {
			// print
			fmt.Println(nowNode.Data)
			// get the next node
			nowNode = nowNode.NextNode
			continue
		}

		// if the next node is null
		break
	}

}
