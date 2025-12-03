package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。

func main() {
	var wg sync.WaitGroup
	var count int64 = 0 // atomic仅支持固定类型（如int64），不能用int

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go plusAtomic(&count, &wg)
	}

	wg.Wait()

	// 原子读取计数器值（保证读取的是最新值）
	finalCount := atomic.LoadInt64(&count)
	fmt.Println("最终计数器值：", finalCount)
}

// 用atomic实现计数器递增
func plusAtomic(counter *int64, wg *sync.WaitGroup) {
	defer wg.Done()
	// 每个协程递增1000次
	for i := 0; i < 10000; i++ {
		// 原子递增：第二个参数是递增步长（+1）
		atomic.AddInt64(counter, 1)
	}
}
