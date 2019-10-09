package main

import (
	"fmt"
	"os"
)

/**
与echo123功能相同,但是将索引和值一同打印了出来
*/
func main() {
	for index, arg := range os.Args[1:] {
		fmt.Printf("index:%d arg:%s\n", index, arg)
	}
}
