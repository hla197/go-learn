package main

import "fmt"

// 题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
func main() {
	nums := []int{1, 2, 3, 4, 5}
	divSlice(&nums)
	fmt.Println(nums)
}

func divSlice(num *[]int) {
	for i := range *num {
		(*num)[i] *= 2
	}
}
