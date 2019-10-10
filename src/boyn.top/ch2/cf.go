package main

import (
	"boyn.top/ch2/tempconv"
	"os"
	"strconv"
)
import "fmt"

/*
这个程序从我们写好的tempconv包导入了一些温度相关的常量,类型和函数
从控制台的执行参数之中读入一个浮点数,计算其作为摄氏度转换为华氏度的值
和计算其作为华氏度转换为摄氏度的值
*/
func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}
