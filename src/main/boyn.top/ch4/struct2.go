package main

import "fmt"

type Point2 struct {
	x, y int
}

/*
如果结构体的全部成员都是可以比较的，
那么结构体也是可以比较的，
那样的话两个结构体将可以使用==或!=运算符进行比较。
相等比较运算符==将比较两个结构体的每个成员
*/
func main() {
	p1 := Point2{1, 2}
	p2 := Point2{y: 2, x: 1}
	fmt.Println(p1 == p2)
}
