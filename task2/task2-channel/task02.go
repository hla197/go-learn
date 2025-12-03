package main

import (
	"fmt"
	"sync"
	"time"
)

// 实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印

func main() {
	ch := make(chan int, 10)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer close(ch)
		num := 1
		for i := 0; i < 10; i++ {
			// 每次10条数据，休眠500毫秒
			for j := 0; j < 10; j++ {
				ch <- num
				num++
			}
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			select {
			case v, ok := <-ch:
				if !ok {
					return
				}
				fmt.Println("接受到数据", v)
			}
		}
	}()

	wg.Wait()
	fmt.Println("end")
}
