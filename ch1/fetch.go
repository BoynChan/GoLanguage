package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

/*
这个程序会获取对应的url,并将其源文本打印出来,如curl命令
*/
func main() {
	for _, url := range os.Args[1:] {
		//如果url没有加上http协议的前缀,则自动为其加上
		if strings.HasPrefix(url, "http://") == false {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch error: %v\n", err)
			os.Exit(1)
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch error: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		//打印出状态码
		fmt.Printf("status code: %d\n", resp.StatusCode)

		//函数调用io.Copy(dst, src)会从src中读取内容，
		// 并将读到的结果写入到dst中，
		// 使用这个函数替代掉例子中的ioutil.ReadAll
		// 来拷贝响应结构体到os.Stdout，避免申请一个缓冲区（例子中的b）来存储

		io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()

		//b,err := ioutil.ReadAll(resp.Body)
		//fmt.Printf("%s",b)
	}
}
