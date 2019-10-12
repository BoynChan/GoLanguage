package main

import "fmt"

func main() {
	x := []int{1, 2, 3}
	fmt.Println(appendInt(x, 4))
}

/*
演示了len与cap的不同
slice中有可能len与cap不同,cap是最大可以支持的容量
当len比cap小的时候,我们向slice中添加元素到len时会扩容
每次调用appendInt函数，
必须先检测slice底层数组是否有足够的容量来保存新添加的元素。
如果有足够空间的话，直接扩展slice（依然在原有的底层数组之上），
将新添加的y元素复制到新扩展的空间，并返回slice。
因此，输入的x和输出的z共享相同的底层数组。
如果没有足够的增长空间的话，
appendInt函数则会先分配一个足够大的slice用于保存新的结果，
先将输入的x复制到新的空间，然后添加y元素。结果z和输入的x引用的将是不同的底层数组。
*/
func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen < cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

/*
剔除字符串slice中为空的元素
*/
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

/*
原地完成消除[]string中相邻重复的字符串
*/
func nonredundant(s []string) []string {
	i := 0
	redundant := ""
	for _, v := range s {
		if redundant == "" {
			redundant = v
		}
		if redundant != v {
			s[i] = v
			i++
			redundant = v
		}
	}
	return s[:i]
}

/*
移除slice中的某个元素
*/
func remove(s []int, position int) []int {
	copy(s[position:], s[position+1:])
	return s[:len(s)-1]
}
