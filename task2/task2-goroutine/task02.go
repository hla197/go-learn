package main

import (
	"fmt"
	"sync"
	"time"
)

// 设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。

type Task struct {
	Name string
	Fc   func()
}

func main() {
	tasks := []Task{
		{
			Name: "task1",
			Fc: func() {
				fmt.Println("task1开始运行")
				time.Sleep(time.Second * 2)
			},
		},
		{
			Name: "task2",
			Fc: func() {
				fmt.Println("task2开始运行")
				time.Sleep(time.Second * 1)
			},
		},
		{
			Name: "task3",
			Fc: func() {
				fmt.Println("task3开始运行")
				time.Sleep(time.Second * 3)
			},
		},
	}

	// 初始化一个等待组（WaitGroup），并设置需要等待的协程数量为任务切片 tasks 的长度，确保主协程会等待所有任务协程执行完成后再退出。
	var wg sync.WaitGroup
	//wg.Add(len(tasks))
	for _, task := range tasks {
		wg.Add(1) // 计算器+1
		go goroutineRun(&wg, task.Name, task.Fc)
	}
	// 阻塞当前协程，直到计数器归 0
	wg.Wait()
	fmt.Println("所有任务执行完毕")

}

func goroutineRun(wg *sync.WaitGroup, name string, fc func()) {
	// 将计数器 -1（协程执行完成时调用）
	defer wg.Done()
	// 记录方法执行前的起始时间
	start := time.Now()

	fc()

	// 计算并输出执行耗时
	elapsed := time.Since(start)

	fmt.Printf("任务[%s]执行耗时：%s\n", name, elapsed)
}
