package main

import (
	"fmt"
	"time"
)

func noDeadLock() {
	c1 := make(chan int)
	go func() {
		c1 <- 1
	}()
	fmt.Println(<-c1)
}

func deadLock() {
	c1 := make(chan int)
	c1 <- 1
	fmt.Println(<-c1)
}

func noBufChan(n int) {
	c1 := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			fmt.Println("send", i)
			c1 <- i
		}
		close(c1)
	}()

	for i := range c1 {
		fmt.Println("receive", i)
	}
}

func bufChan(n int) {
	c1 := make(chan int, 2)
	go func() {
		for i := 0; i < n; i++ {
			fmt.Println("send", i)
			c1 <- i
		}
		close(c1)
	}()

	for i := range c1 {
		fmt.Println("receive", i)
	}
}

func bufChanEquivalent(n int) {
	c1 := make(chan int, 2)
	go func() {
		for i := 0; i < n; i++ {
			fmt.Println("send", i)
			c1 <- i
		}
		close(c1)
	}()

	// for range c1 的等价写法
	for {
		i, ok := <-c1
		if !ok {
			break // channel已关闭，退出循环
		}
		fmt.Println("receive", i)
	}
}

func noBufChanEquivalent(n int) {
	c1 := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			fmt.Println("send", i)
			c1 <- i
		}
		close(c1)
	}()

	// for range c1 的等价写法
	for {
		i, ok := <-c1
		if !ok {
			break // channel已关闭，退出循环
		}
		fmt.Println("receive", i)
	}
}

func bufChanForNumber(n int) {
	c1 := make(chan int, 2)
	go func() {
		for i := 0; i < n; i++ {
			fmt.Println("send", i)
			c1 <- i
		}
		close(c1)
	}()

	for i := 0; i < n; i++ {
		fmt.Println("receive", <-c1)
	}
}

func noBufChanForNumber(n int) {
	c1 := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			fmt.Println("send", i)
			c1 <- i
		}
	}()

	for i := 0; i < n; i++ {
		fmt.Println("receive", <-c1)
	}
}

func noBufChanSync(n int) {
	c1 := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < n; i++ {
			fmt.Println("send", i)
			c1 <- i
			<-done // 等待接收方完成
		}
	}()

	for i := 0; i < n; i++ {
		fmt.Println("receive", <-c1)
		done <- true // 通知发送方继续
	}
}

func selectChan(n int) {
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)

	go func() {
		for i := 0; i < n; i++ {
			fmt.Println("send c1", i)
			c1 <- i
			fmt.Println("send c2", i)
			c2 <- i
			fmt.Println("send c3", i)
			c3 <- i
		}
		// close(c1)
		// close(c2)
		// close(c3)
	}()

	// 1. 基本用法：从多个channel中随机选择一个可用的
	for i := 0; i < n*3; i++ {
		select {
		case i := <-c1:
			fmt.Println("c1 receive", i)
		case i := <-c2:
			fmt.Println("c2 receive", i)
		case i := <-c3:
			fmt.Println("c3 receive", i)
		}
	}
}

// 2. 带超时的select
func selectWithTimeout() {
	c := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		c <- 1
	}()

	select {
	case i := <-c:
		fmt.Println("received", i)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	}
}

// 3. 带default的select（非阻塞）
func selectWithDefault() {
	c := make(chan int)
	select {
	case i := <-c:
		fmt.Println("received", i)
	default:
		fmt.Println("no data available")
	}
}

// 4. 关闭channel的检测
func selectWithClose() {
	c := make(chan int)
	go func() {
		close(c)
	}()

	select {
	case i, ok := <-c:
		if !ok {
			fmt.Println("channel closed")
			return
		}
		fmt.Println("received", i)
	}
}

// 5. 发送和接收的select
func selectSendReceive() {
	c1 := make(chan int)
	c2 := make(chan int)
	go func() {
		c1 <- 1
	}()

	select {
	case i := <-c1:
		fmt.Println("received", i)
	case c2 <- 2:
		fmt.Println("sent 2")
	}
}

func selectDeadlock(n int) {
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)

	go func() {
		for i := 0; i < n; i++ {
			fmt.Println("send c1", i)
			c1 <- i
			fmt.Println("send c2", i)
			c2 <- i
			fmt.Println("send c3", i)
			c3 <- i
		}
		// close(c1)
		// close(c2)
		// close(c3)
	}()

	// select {
	// case i := <-c1:
	// 	fmt.Println("c1 receive", i)
	// case i := <-c2:
	// 	fmt.Println("c2 receive", i)
	// case i := <-c3:
	// 	fmt.Println("c3 receive", i)
	// 	// default:
	// 	// 	fmt.Println("default")
	// }

	// 如果 循环次数 大于 n*3 ，则会导致死锁
	// 等待所有数据
	for i := 0; i < n*3; i++ {
		select {
		case i := <-c1:
			fmt.Println("c1 receive", i)
		case i := <-c2:
			fmt.Println("c2 receive", i)
		case i := <-c3:
			fmt.Println("c3 receive", i)
		}
	}
}

func readWriteChan(n int) {
	c1 := make(chan int)
	var readc <-chan int = c1
	var writec chan<- int = c1

	writeChan := func() {
		for i := 0; i < n; i++ {
			fmt.Println("write", i)
			writec <- i
		}
		close(writec)
	}

	readChan := func() {
		for i := 0; i < n; i++ {
			fmt.Println("read", <-readc)
		}
	}

	go writeChan()
	go readChan()
}

func main() {
	n := 5

	noDeadLock()
	fmt.Println("=== 无缓冲channel ===")
	noBufChan(n)
	fmt.Println("=== 有缓冲channel ===")
	bufChan(n)
	fmt.Println("=== 有缓冲channel(等价写法) ===")
	bufChanEquivalent(n)
	fmt.Println("=== 无缓冲channel(等价写法) ===")
	noBufChanEquivalent(n)
	fmt.Println("=== 无缓冲channel(for Number循环) ===")
	noBufChanForNumber(n)
	fmt.Println("=== 无缓冲channel(同步发送接收) ===")
	noBufChanSync(n)
	fmt.Println("=== select channel ===")
	selectChan(n)
	fmt.Println("=== select with timeout ===")
	selectWithTimeout()
	fmt.Println("=== select with default ===")
	selectWithDefault()
	fmt.Println("=== select with close ===")
	selectWithClose()
	fmt.Println("=== select send receive ===")
	selectSendReceive()
	fmt.Println("=== select deadlock ===")
	selectDeadlock(n) // 增大里面select循环次数，会导致死锁
	fmt.Println("=== read write chan ===")
	readWriteChan(n)
}
