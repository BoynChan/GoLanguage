/*
我们将看到如何使用Go语言提供的不同寻常的结构体嵌入机制让一个
命名的结构体包含另一个结构体类型的匿名成员，
这样就可以通过简单的点运算符x.f来访问匿名成员链中嵌套的x.d.e.f成员
*/
package main

import "fmt"

/*
考虑一个二维的绘图程序，提供了一个各种图形的库，
例如矩形、椭圆形、星形和轮形等几何形状。这里是其中两个的定义：
*/
/*
Go语言有一个特性让我们只声明一个成员对应的数据类型而不指名成员的名字；
这类成员就叫匿名成员。匿名成员的数据类型必须是命名的类型或指向一个命名的类型的指针
*/
type Point struct {
	X, Y int
}
type Circle struct {
	Point
	Radius int
}
type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var w Wheel
	w.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20
	w2 := Wheel{
		Circle: Circle{
			Point: Point{
				X: 1,
				Y: 2,
			},
			Radius: 5,
		},
		Spokes: 20,
	}
	//需要注意的是Printf函数中%v参数包含的#副词，
	// 它表示用和Go语言类似的语法打印值。对于结构体类型来说，将包含每个成员的名字
	fmt.Printf("%#v\n", w2)
}
