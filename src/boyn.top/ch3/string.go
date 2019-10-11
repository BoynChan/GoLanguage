package main

import (
	"bytes"
	"fmt"
	"strings"
)

/*
下面这个例子的basename函数与unix系统的basename相同
将系统路径的前缀删除,将文件类型的后缀名删除
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
*/
func basename1(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			fmt.Println(s)
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
		}
	}
	return s
}

func basename2(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

/*
函数的功能是将一个表示整值的字符串，每隔三个字符插入一个逗号分隔符，例如“12345”处理后成为“12,345”。这个版本只适用于整数类型
*/
func comma(s string) string {
	n := len(s)
	if n < 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func commaWithFloat(s string, isPositive bool) string {
	var buf bytes.Buffer
	if isPositive {
		buf.WriteString("+")
	} else {
		buf.WriteString("-")
	}
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		buf.WriteString(comma(s[:dot]))
		buf.WriteString(s[dot:])
		return buf.String()
	} else {
		buf.WriteString(comma(s))
		return buf.String()
	}
}

func commaWithBuffer(s string) string {
	var buf bytes.Buffer
	if len(s) <= 3 {
		buf.WriteString(s)
		return buf.String()
	}
	writeCommaLength := len(s) % 3
	buf.WriteString(s[0:writeCommaLength])
	i := writeCommaLength
	for i < len(s) {
		if i != 0 {
			buf.WriteString(",")
		}
		buf.WriteString(s[i : i+3])
		i += 3
	}
	return buf.String()
}

func intsToStrings(values []int) string {
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range values {
		if i > 0 {
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteString("]")
	return buf.String()
}

func main() {
	/*	fmt.Println(basename1("a.go"))
		fmt.Println(basename1("/a/b/c.go"))
		fmt.Println(basename1("abc.d.go"))*/
	fmt.Println(basename2("/a/b/c.go"))
	fmt.Println(comma("12345"))
	fmt.Println(commaWithBuffer("1234567899"))
	fmt.Println(commaWithFloat("1234567899.1234", true))
	fmt.Println(intsToStrings([]int{1, 2, 3, 4, 5}))
}
