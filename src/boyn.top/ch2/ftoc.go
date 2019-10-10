package main

import "fmt"

/*
这个函数与boiling.go所实现的功能基本相同
他们的区别在于,当前文件将华氏度转为摄氏度这个逻辑提取了出来
转变为一个函数,这样的话这个函数就可以在多个地方被多次使用
*/
func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("f=%g,c=%g", freezingF, fToC(freezingF))
	fmt.Printf("f=%g,c=%g", boilingF, fToC(boilingF))
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
