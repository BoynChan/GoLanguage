package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

/*
演示了数组的另外一种初始化的办法,就是可以指定一个索引和对应值列表的方式初始化
*/
func main() {
	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	//symbol2 := [...]string{0: "$", 1: "€", 2: "￡", 3: "￥",} 上下两者等价
	fmt.Printf("the symbol of RMB is %s", symbol[RMB])
}
