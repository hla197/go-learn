package main

import (
	"fmt"
	"sync"
)

// 编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup
	count := 0

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go plus(&mu, &wg, &count)
	}

	wg.Wait()
	fmt.Println("count", count)
}

func plus(mu *sync.Mutex, wg *sync.WaitGroup, counter *int) {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		mu.Lock()
		*counter++
		mu.Unlock()
	}
}
