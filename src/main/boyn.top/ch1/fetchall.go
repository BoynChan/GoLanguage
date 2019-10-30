package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

/*
这个程序使用了go语言中的强大特性goroutine和channel来进行并发地请求多个地址
结果如下
1.42s     243 http://httpbin.org/get?a=2
1.46s     243 http://httpbin.org/get?a=3
1.50s     243 http://httpbin.org/get?a=1
1.51s     243 http://httpbin.org/get?a=4
pass time 1.51s
可以看到,每一个地址都会用1.4~1.5秒,但是由于其是并发请求的,所以请求的时间不会超过最长的一个的请求时间
*/
func main() {
	start := time.Now()
	ch := make(chan string) //一个传递string类型参数的channel
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) //从通道ch中读取
	}
	fmt.Printf("pass time %.2fs", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	if strings.HasPrefix(url, "http://") == false {
		url = "http://" + url
	}
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	//可以把Discard这个变量看作一个垃圾桶，可以向里面写一些不需要的数据
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s:%v", url, err)
		os.Exit(1)
	}
	//从start所指示的时间开始经过了多少时间
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
