package main

import (
	"fmt"
	"sync"
)

// 题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。

func main() {
	// 定义等待组，等待两个协程执行完成
	var wg sync.WaitGroup
	wg.Add(2) // 计数+2，对应两个协程

	go printOdd(&wg)
	go printEven(&wg)

	wg.Wait()
	fmt.Println("所有协程执行完成")
}

func printOdd(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i += 2 {
		fmt.Println(i)
	}
}

func printEven(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		fmt.Println(i)
	}
}
