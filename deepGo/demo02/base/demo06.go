// @Author Gopher
// @Date 2025/1/19 21:14:00
// @Desc
package main

import "fmt"

// Speaker 接口要求实现者必须提供 Speak() 方法，该方法返回一个字符串。
type Speaker interface {
	Speak() string
}

// Greeter 接口要求实现者必须提供 Greet(name string) 方法，该方法接受一个字符串作为参数，并返回一个字符串。
type Greeter interface {
	Greet(name string) string
}

// Human 是一个结构体，它包含一个字段 Name，用来存储人的名字。这个结构体没有实现接口，但可以通过实现接口所需的方法来使 Human 类型满足接口的要求。
type Human struct {
	Name string
}

// Speak
/**
Human 类型实现了 Speak() 和 Greet(name string) 两个方法：
Speak() 方法返回一个包含 Name 字段的问候语，例如 Hello, my name is Alice。
Greet(name string) 方法返回一个打招呼的句子，例如 Hello, Alice。
*/
func (h *Human) Speak() string {
	return "Hello, my name is " + h.Name
}

// Greet 注意这里使用了 *Human（Human 的指针类型）作为方法的接收者。这样，方法能够修改 Human 类型的值，并且允许通过指针类型的变量调用这些方法。
func (h *Human) Greet(name string) string {
	return "Hello, " + name
}

func main() {

	var speaker Speaker = &Human{Name: "Alice"}
	fmt.Println(speaker.Speak()) // 输出: Hello, my name is Alice

	var greeter Greeter = &Human{Name: "Bob"}
	fmt.Println(greeter.Greet("Alice")) // 输出: Hello, Alice

}
