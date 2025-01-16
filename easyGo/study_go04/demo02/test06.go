package main

import (
	"fmt"
	"strconv"
)

type goods struct {
	name  []string
	price []int
	num   []int
}

func (g *goods) show() string {
	var info string
	for k, v := range g.name {
		info += "商品名称：" + v
		info += "商品单价：" + strconv.Itoa(g.price[k]) + "元"
		info += "商品数量：" + strconv.Itoa(g.num[k]) + "\n"
	}
	return info
}

func (g *goods) init(name []string, price []int, num []int) {
	g.name = name
	g.price = price
	g.num = num
}

func main() {

	g := &goods{}
	g.init([]string{"品牌手机", "品牌计算机"}, []int{1699, 2399}, []int{2, 3})
	info := g.show()
	fmt.Println(info)
}
