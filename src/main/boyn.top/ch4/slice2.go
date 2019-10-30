package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(array)
	//reverseByPointer(&array)
	reverse(array)
	fmt.Println(array)
	array = []int{1, 2, 3, 4, 5, 6, 7}
	rotateLeft(array, 2)
	fmt.Println(array)
	array = []int{1, 2, 3, 4, 5, 6, 7}
	rotateRight(array, 2)
	fmt.Println(array)
}

/*
利用切片的特性来倒转数组
*/
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

/*
传入数组指针来实现倒转数组
*/
func reverseByPointer(s *[]int) {
	len := len(*s)
	for i, j := 0, len-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

/*
利用切片和反转的特性实现向左旋转数组
先倒转前面n位
再倒转后面len-n位
最后整体倒转
*/
func rotateLeft(s []int, n int) {
	reverse(s[0:n])
	reverse(s[n:])
	reverse(s)
}

/*
利用切片和反转的特性实现向右旋转数组
先整体倒转
再倒转前面n位
最后倒转后面len-n位
*/
func rotateRight(s []int, n int) {
	reverse(s)
	reverse(s[0:n])
	reverse(s[n:])
}

func rotate(s []int) {

}

/*
比较两个slice是否相同
*/
func equals(src1 []int, src2 []int) bool {
	if len(src1) != len(src2) {
		return false

	}
	for i := range src1 {
		if src1[i] != src2[i] {
			return false
		}
	}
	return true

}
