package main

import "fmt"

/*
slice与python中的数组切片类似,是一个很方便的操作可变数组的工具
slice底层也是由数组来表示的
*/
func main() {
	slice := [...]int{1, 2, 3, 4}
	fmt.Println(slice[0:])
	fmt.Println(slice[:3])
}
