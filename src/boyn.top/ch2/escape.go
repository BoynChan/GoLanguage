package main

var global *int

/*
在这个例子中,虽然x是在函数内部定义的,但是因为外部变量对x进行了引用,所以他的作用域被扩大到包一级的范围
我们称为x变量从函数中逃逸了
*/
func main() {
	x := 1
	global = &x
}
