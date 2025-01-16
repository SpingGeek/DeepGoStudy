package main // 声明 main 包

import "fmt" // 导入 fmt 包，用于打印字符串

type carInMotion interface { // 为行驶中的汽车声明接口
	move(speed int)                                                           // 行驶函数。参数是汽车此时行驶的速度
	brake() (int, int)                                                        // 刹车函数。返回值分别是刹车前的速度和刹车后的速度
	park()                                                                    // 泊车函数
	consumeOil(fuelLeft float64, aver_consumption float64) (distance float64) // 消耗燃油函数，返回行驶的距离
}

type car struct { // 表示汽车的结构体
	color string // 汽车的颜色
}

func (c *car) move(speed int) { // 实现接口中的 move
	fmt.Printf("一辆%v 的汽车在以 %v km/h 的速度匀速行驶\n", c.color, speed)
}

func (c *car) brake() (int, int) {
	fmt.Printf("这辆 %v 的汽车开始刹车\n", c.color)
	speedBeforeBrake := 60
	speedAfterBrake := 0
	return speedBeforeBrake, speedAfterBrake
}

func (c *car) park() {
	fmt.Printf("这辆 %v 的汽车停在路边的车位里面\n", c.color)
}

func (c *car) consumeOil(fuelLeft float64, aver_consumption float64) (distance float64) {
	fmt.Printf("这辆 %v 的汽车剩余燃油 %vL, 燃油的平均油耗 %v L/100km\n", c.color, fuelLeft, aver_consumption)
	return fuelLeft / aver_consumption * 100
}

func main() {
	var cim carInMotion
	cr := car{color: "红色"}
	cim = &cr
	cim.move(60)
	speedBeforeBrake, speedAfterBrake := cim.brake()
	fmt.Printf("before the brake %v km/h, after the brake %v km/h", speedBeforeBrake, speedAfterBrake)
	cim.park()
	distance := cim.consumeOil(27, 6.3)
	fmt.Printf("continue to drive %.2f km\n", distance)

}
