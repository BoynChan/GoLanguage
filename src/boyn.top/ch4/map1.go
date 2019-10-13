package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	wordfreq()
}

/*
创建一个map
*/
func createMap() {
	ages := make(map[string]int)
	ages["joy"] = 1
	ages["fun"] = 2
}

/*
在map中,如果没有找到对应元素,则会返回值类型的对应零值,所以下面行为是安全的
*/
func safeAdd() {
	ages := make(map[string]int)
	ages["anne"] = ages["anne"] + 1
	//在执行后,anne键的值为1
}

/*
删除map中的元素
*/
func deleteElement() {
	ages := make(map[string]int)
	ages["joy"] = 1
	ages["fun"] = 2
	delete(ages, "joy")
}

/*
迭代循环每一个值
*/
func iterator() {
	ages := make(map[string]int)
	ages["joy"] = 1
	ages["fun"] = 2
	for key, value := range ages {
		fmt.Printf("name:%s,age:%d\n", key, value)
	}
}

/*
可以通过令map的取值操作返回两个值来确定某个索引值是否存在
*/
func getIf() {
	ages := make(map[string]int)
	ages["joy"] = 1
	ages["fun"] = 2
	if value, ok := ages["abc"]; !ok {
		fmt.Println("Not exist")
	} else {
		fmt.Println(value)
	}
}

/*
下面的dedup程序读取多行输入，但是只打印第一次出现的行
主要原理就是将值类型设置为布尔型,当第一次输入时,为布尔值的空值也就是false
这时候将其添加到map中,然后再有同样类型的值就不会添加
*/
func dedup() {
	seen := make(map[string]bool) // a set of strings
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}

/*
ReadRune方法执行UTF-8解码并返回三个值：
解码的rune字符的值，字符UTF-8编码后的长度，和一个错误值。
我们可预期的错误值只有对应文件结尾的io.EOF。
如果输入的是无效的UTF-8编码的字符，
返回的将是unicode.ReplacementChar表示无效字符，并且编码长度是1。

charcount程序同时打印不同UTF-8编码长度的字符数目。
对此，map并不是一个合适的数据结构；
因为UTF-8编码的长度总是从1到utf8.UTFMax（最大是4个字节），使用数组将更有效。
*/
func charcount() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount:%v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for char, count := range counts {
		fmt.Printf("%q\t%d\n", char, count)
	}
	fmt.Printf("length\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

/*
统计单词频率
*/
func wordfreq() {
	counts := make(map[string]int)
	in := bufio.NewReader(os.Stdin)
	for {
		word, _, err := in.ReadLine()
		if err != nil {
			break
		}
		words := strings.Fields(string(word))
		for _, v := range words {
			counts[v]++
		}
	}
	fmt.Printf("word\tcount\n")
	for word, count := range counts {
		fmt.Printf("%s\t%d\n", word, count)
	}
}
