package main

import "fmt"

//var const type func 分别对应变量,常量,类型和函数的声明

const boilingF = 212.0 //这是一个常量,它在包一级范围声明,这个常量可以在相同包下的每一个源文件中进行访问.

func main() { //main是一个函数
	//:=是一种简短变量声明的表示,它不需要声明变量理性,而是会根据表达式来自动推导
	//大部分局部变量的声明和初始化都可以使用这个表示
	// =是赋值 :=是声明  记得区分
	f := boilingF         //f是一个变量
	c := (f - 32) * 5 / 9 //c也是一个变量,在很多时候,声明变量时不需要显式声明它的类型,而是会进行自动推断
	fmt.Printf("f=%g, c=%g\n", f, c)
}