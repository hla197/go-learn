package main

import (
	"fmt"
	"sync"
	"time"
)

// 编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。

func main() {
	ch := make(chan int, 5)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer close(ch)
		for i := 1; i <= 10; i++ {
			ch <- i
			fmt.Println("input", i)
		}
	}()
	timeout := time.After(time.Second * 1)
	go func() {
		defer wg.Done()
		for {

			select {
			case v, ok := <-ch:
				if !ok {
					fmt.Println("Channel已关闭")
					return
				}
				fmt.Printf("get: %d\n", v)
			case <-timeout:
				fmt.Println("超时")
				return
			default:
				fmt.Println("等待数据中...")
				time.Sleep(500 * time.Millisecond)
			}

		}
	}()
	// 阻塞等待完成
	wg.Wait()
	fmt.Println("end")
}
