package main

import "fmt"

func incr(v *int) int {
	*v++
	return *v
}

/*
这个文件演示了指针的使用,在go语言中指针的使用与c语言中的使用非常类似
&是取地址符
*是取值符
*/
func main() {
	v := 1
	p := &v
	incr(p)

	fmt.Printf("now v is %d", incr(p))
}
