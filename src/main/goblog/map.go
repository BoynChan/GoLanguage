package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	//创建了一个映射类型,键为string,值为int
	words := make(map[string]int)
	//strings.Fields用来将语句按照语序分开,比如空格,逗号等等
	for _, v := range strings.Fields(s) {
		//当key在这个映射中不存在,默认返回该值类型的默认零值
		//在这里就是0了
		words[v] = words[v] + 1
	}
	//delete(words,key) 删除元素
	return words
}

func main() {
	fmt.Println(WordCount("The quick,brown fox jumped over the lazy dog."))
}
