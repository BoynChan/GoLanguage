package main

import "fmt"

func main() {
	//这样会创建一个数组
	//数组类型长度和容量是对等的,并且是在定义时就已经确定好了.无法变化
	//数组的访问方式跟c一样
	var a [4]int
	a[0] = 5

	b := [2]string{"Penn", "Teller"}
	b[0] = "Go"
	//b[2] = "Go" 这样会出错,因为超过了数组的容量

	//切片在创建时可以不指定容量,并且可以一直增加,它会自动扩容
	//创建了一个长度为5,容量为5的切片
	letter := []string{"a", "b", "c"}
	letter = make([]string, 5, 5)
	//用append函数来进行元素的增加
	letter = append(letter, "d")
	//切片十分有用的作用之一,就是可以像这样 任意选取元素,:前后可以省略,默认为0和len(slice)
	letter = letter[2:]
	letter = letter[:]

	//增加容量的方法1,创建一个新的切片,将容量设置得大一些
	//然后复制源数组,并将新的数组的引用转到源数组
	t := make([]string, len(letter), ((cap(letter) + 1) * 2))
	copy(t, letter)
	letter = t

	//增加容量的方法2,append方法
	//这是一个内置的方法,对切片进行动态扩容
	//我们只需要指定好增加的元素即可,因为其内部会动态扩容
	// num = [0 1 2 3 4 5 6]
	num := make([]int, 1, 1)
	num = append(num, 1, 2, 3, 4, 5, 6)

	// 可以这样来复制一个切片
	// num2 = [0 0 1 2 3 4 5 6]
	num2 := make([]int, 0)
	num2 = append(num2, num...)
	//用这种方式来遍历切片
	for i, v := range num2 {
		num2[i] = v
	}
	fmt.Println(num2)

	// 对于数组来说,不能像切片一样被很方便地拓展
	// 所以需要通过一个显式的循环来进行复制
	// num3 = [Go Teller]
	num3 := make([]string, 0)
	for _, v := range b {
		num3 = append(num3, v)
	}
	fmt.Println(num3)

}
