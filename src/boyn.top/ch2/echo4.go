package main

import (
	"flag"
	"fmt"
	"strings"
)

//flag的Bool,String都会生成一个对应类型的指针
var n = flag.Bool("n", false, "omit trailing newline")
var seq = flag.String("s", " ", "separator")

/*
这个程序使用了flag库,他会根据我们的声明,接收对应的参数并做出一定的处理
*/
func main() {
	flag.Parse()
	//如果没有指定-s 后面的参数,则默认以空格为分隔符
	fmt.Printf(strings.Join(flag.Args(), *seq))
	//如果没有指定n,则默认在输出后进行换行
	if !*n {
		fmt.Println()
	}
}
