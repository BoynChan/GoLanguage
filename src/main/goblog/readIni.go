package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getValue(filename string, expectSection string, expectKey string) string {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Cannot find file")
	}
	//defer是一个延迟执行的关键字,会在函数退出的时候调用
	defer file.Close()

	reader := bufio.NewReader(file)

	var sectionName string

	for {
		linestr, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		linestr = strings.TrimSpace(linestr)
		if linestr == "" {
			continue
		}

		if linestr[0] == ';' {
			continue
		}

		if linestr[0] == '[' && linestr[len(linestr)-1] == ']' {
			sectionName = linestr[1 : len(linestr)-1]
		} else if sectionName == expectSection {
			pair := strings.Split(linestr, "=")
			if len(pair) == 2 {
				key := strings.TrimSpace(pair[0])
				if key == expectKey {
					return strings.TrimSpace(pair[1])
				}
			}
		}
	}
	return ""
}

func main() {
	filename := os.Args[1]
	section := os.Args[2]
	key := os.Args[3]

	value := getValue(filename, section, key)
	fmt.Println(value)
}
