package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
功能与dup1相同,但是这里加入了一个读取文件中数据的办法
可以通过参数指定文件名字,进而打开文件内容.
*/
func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			//如果err对象为空,则证明文件成功打开
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "dup2: err %s\n", err)
				continue
			}
			//传入文件对象,并且在读取完成后关闭文件流
			countLines(f, counts)
			_ = f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%s show %d times", line, n)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
