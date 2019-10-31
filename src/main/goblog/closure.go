package main

import "fmt"

//使用了函数的闭包特性
//counter对于新创建的counter函数中返回的匿名函数来说,
//是一个隐式的变量,可以被函数引用但是外界无法引用
func counter(init int) func() int {
	counter := init
	return func() int{
		counter++
		return counter
	}
}

//这是一个利用闭包特性
//完成斐波那契数列的函数
//引用的变量有front,last和n
//用n来记录引用到第n个元素
//front和last用来算当前的元素
//我们知道斐波那契数列的每一个元素是前两个元素之和,所以我们把前两个元素加起来,并更新这两个元素
//达到高效计算斐波那契数列的效果
func fibonacci() func() int {
	front := 0;
	last := 1;
	n := 0;
	return func() int {
		if n == 0 {
			n++
			return 0
		}else if n == 1{
			n++
			return 1
		}else {
			temp := last
			last = front + last
			front = temp
			n++
			return last
		}
	}
}

func main() {
	incr := counter(0)
	fmt.Println(incr())
	fmt.Println(incr())
}

