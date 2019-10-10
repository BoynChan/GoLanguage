package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

/**
这个程序调用了io工具库 ioutil中的ReadFile函数,它会将文件全部读入,而不是以流的形式读入
我们将全部数据读入后,将其转成string,再以换行符作为分隔符将其分开
ReadFile返回的是一个字节类型的slice,必须要被转成string才可以被分割
实际上,无论是这个程序中的ioutil还是之前的bufio,都是对os.File中的Read和Write函数的封装
不过我们很少会直接使用这两个函数,而是使用上面提到的一些更加易用的官方库
*/
func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "openfile error: %s", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		fmt.Printf("%d  %s \n", n, line)
		if n > 1 {
			fmt.Printf("%d  %s \n", n, line)
		}
	}
}
