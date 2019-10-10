package main

/**
与echo1,2功能相同,但是使用了strings的Join功能将字符串拼接在一起
*/
import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}
