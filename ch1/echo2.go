package main

/**
与echo1功能相同,但是使用了for-each迭代
*/
import (
	"fmt"
	"os"
)

func main() {
	s, seq := "", ""
	//由于go不允许出现未使用的变量,所以如果我们在进行range迭代的时候不需要使用索引值,可以用一个下划线符号(_)进行代替
	//这个符号通常用于语法上需要使用的变量名但是程序中不需要
	for _, arg := range os.Args[1:] {
		s += seq + arg
		seq = " "
	}
	fmt.Println(s)

}
