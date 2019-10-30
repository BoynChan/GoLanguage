package main

import (
	"crypto/sha256"
	"fmt"
)

/*
crypto/sha256包的Sum256函数对一个任意的字节slice类型的数据生成一个对应的消息摘要。
消息摘要有256bit大小，因此对应[32]byte数组类型。
如果两个消息摘要是相同的，那么可以认为两个消息本身也是相同
（译注：理论上有HASH码碰撞的情况，但是实际应用可以基本忽略）；
如果消息摘要不同，那么消息本身必然也是不同的。下面的例子用SHA256算法分别生成“x”和“X”两个信息的摘要
*/
func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%x\n", c1, c2, zero(&c1))
	fmt.Printf("%d\n", diff(&c1, &c2))
}

/*
比较两个数组中不同的字节数
*/
func diff(ptr1 *[32]byte, ptr2 *[32]byte) int {
	count := 0
	for i := range ptr1 {
		if ptr1[i] != ptr2[i] {
			count++
		}
	}
	return count
}

/*
对于数组,也可以传入一个数组指针
这样对数组的修改就可以直接反馈到本数组中
*/
func zero(ptr *[32]byte) [32]byte {
	for i := range ptr {
		ptr[i] = 0
	}
	return *ptr
}
