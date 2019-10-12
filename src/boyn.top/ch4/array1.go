package main

import "fmt"

/*
演示数组的基本用法
*/
func main() {
	var array1 = []int{1, 2, 3}
	for i, v := range array1 {
		fmt.Printf("index:%d,value:%d\n", i, v)
	}

	array2 := [...]int{1, 2, 3, 4}
	for i, v := range array2 {
		fmt.Printf("index:%d,value:%d\n", i, v)
	}
}
