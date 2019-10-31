package main

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
	letter := []string{"a","b","c"}
	//创建了一个长度为5,容量为5的切片
	//
	letter = make([]string,5,5)
	letter = append(letter, "d")
	letter = letter[2:]
}