package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	//读取字符长度
	fmt.Println(len("忍者"))
	//读取utf8字符长度
	fmt.Println(utf8.RuneCountInString("忍者"))

	words := "网站:boyn.top"

	//按字符遍历字符串
	for i := 0; i < len(words); i++ {
		fmt.Printf("%c %d ", words[i], words[i])
	}

	fmt.Println()

	//按UTF8遍历字符串
	for _, v := range words {
		fmt.Printf("%c %d ", v, v)
	}

	fmt.Println()

	//Index函数索引字符串
	yinghao := strings.Index(words, ":")
	net := strings.Index(words[yinghao:], "boyn")
	fmt.Println(yinghao, net, words[yinghao+net:])

	host := "localhost"
	port := "8080"
	//声明字节缓冲对象
	var stringBuilder bytes.Buffer

	//把字符串写入缓冲
	stringBuilder.WriteString(host)
	stringBuilder.WriteString(port)
	//输出
	fmt.Println(stringBuilder.String())

	//字符串格式化
	profile := &struct {
		Name string
		Hp   string
	}{
		Name: "rat",
		Hp:   "150",
	}

	fmt.Printf("Use %%+v  %+v \n", profile)
	fmt.Printf("Use %%#v  %#v \n", profile)
	fmt.Printf("Use %%T %T \n", profile)

	//base64编码实例
	message := "Always. --斯内普"
	encodedMessage := base64.StdEncoding.EncodeToString([]byte(message))
	fmt.Println(encodedMessage)
	decodedMessage, err := base64.StdEncoding.DecodeString(encodedMessage)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(decodedMessage))
}
