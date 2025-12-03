package main

import "fmt"

// 定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。

type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
}

func (obj *Rectangle) Area() {
	fmt.Println("this is Rectangle Area()")
}

func (obj *Rectangle) Perimeter() {
	fmt.Println("this is Rectangle Perimeter()")
}

type Circle struct {
}

func (obj *Circle) Area() {
	fmt.Println("this is Circle Area()")
}

func (obj *Circle) Perimeter() {
	fmt.Println("this is Circle Perimeter()")
}
func main() {
	var obj1 Shape = &Rectangle{}
	var obj2 Shape = &Circle{}
	obj1.Area()
	obj1.Perimeter()
	obj2.Area()
	obj2.Perimeter()
}
