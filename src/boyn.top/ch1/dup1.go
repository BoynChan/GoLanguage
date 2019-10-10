package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
功能是读取用户的输入,并找出输入中重复的值
*/
func main() {
	//map是go语言中内建的数据结构,类似于Java中的HashMap
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
