package main

/* Basic Usage:
GoLanguage\src\geek> go run .\accept_parameter.go -name="boyn"
Hello boyn
GoLanguage\src\geek> go run .\accept_parameter.go
Hello everyone
GoLanguage\src\geek> go run .\accept_parameter.go --help
Usage of C:\Users\10655\AppData\Local\Temp\go-build307849546\b001\exe\accept_parameter.exe:
  -name string
        The greeting Object (default "everyone")
exit status 2
*/
import (
	"flag"
)

var name string

func init() {
	//第一个参数用于指定接收结果的变量,第二个是指参数的名字,第三个是默认值,第四个是对这个值的说明
	flag.StringVar(&name, "name", "everyone", "The greeting Object")
}

func main() {
	flag.Parse()
	hello(name)
}
