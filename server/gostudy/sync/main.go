package main

// Mutext 互斥锁

// RWMutex 读写互斥锁
// Mutext 和 RWMutex 的区别
// 1. Mutext 只能有一个写锁，RWMutex 可以有多个读锁
// 2. Mutext 方法有 Lock、Unlock，RWMutex 方法有 Lock、Unlock、RLock、RUnlock

// Once 一次性执行
// 1. Once 方法有 Do，Once 方法可以保证某个函数只执行一次

// WaitGroup 等待组
// 1. WaitGroup 方法有 Add、Done、Wait

// Cond 条件变量
// 1. Cond 方法有 Wait、Signal、Broadcast

// Map 并发安全的map
// 1. Map 方法有 Load、Store、Delete、Range

// Pool 对象池
// 1. Pool 方法有 New

// Semaphore 信号量
// 1. Semaphore 方法有 New

import (
	"fmt"
	"sync"
	"time"
)

// 全局计时器
var startTime time.Time

// 包装函数
func printWithTime(format string, args ...interface{}) {
	elapsed := time.Since(startTime)
	if format == "" {
		fmt.Printf("%v @%.9f\n", args[0], elapsed.Seconds())
	} else {
		fmt.Printf(format+" @%.9f\n", append(args, elapsed.Seconds())...)
	}
}

func testMutex() {
	workerFunc := func(mu *sync.Mutex, wg *sync.WaitGroup) {
		defer wg.Done() // 确保在函数结束时调用 Done
		mu.Lock()
		defer mu.Unlock()
		printWithTime("workerFunc")
		time.Sleep(time.Second)
	}

	mu := &sync.Mutex{}
	var wg sync.WaitGroup // 创建 WaitGroup

	// 启动 5 个 goroutine
	for i := 0; i < 4; i++ {
		wg.Add(1) // 在启动 goroutine 前增加计数
		go workerFunc(mu, &wg)
	}

	// 等待所有 goroutine 执行完毕
	wg.Wait()
	printWithTime("All goroutines completed")
}

func testRWMutex() {
	// 定义一个共享的数据
	var sharedData = 0

	// 读取数据的函数 - 使用读锁
	readerFunc := func(id int, mu *sync.RWMutex, wg *sync.WaitGroup) {
		defer wg.Done()

		// 获取读锁
		mu.RLock()
		defer mu.RUnlock()

		// 读取数据并打印
		printWithTime("Reader %d: reading data: %d", id, sharedData)
		time.Sleep(500 * time.Millisecond) // 模拟读取操作
		printWithTime("Reader %d: finished reading", id)
	}

	// 写入数据的函数 - 使用写锁
	writerFunc := func(id int, newValue int, mu *sync.RWMutex, wg *sync.WaitGroup) {
		defer wg.Done()

		// 获取写锁
		mu.Lock()
		defer mu.Unlock()

		// 写入数据
		printWithTime("Writer %d: writing data, changing from %d to %d", id, sharedData, newValue)
		time.Sleep(1 * time.Second) // 模拟写入操作需要更长时间
		sharedData = newValue
		printWithTime("Writer %d: finished writing", id)
	}

	mu := &sync.RWMutex{}
	var wg sync.WaitGroup

	// 启动多个读取 goroutine - 它们可以并发执行
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go readerFunc(i, mu, &wg)
	}

	// 启动两个写入 goroutine - 它们会互斥执行，并且会阻塞所有读取操作
	wg.Add(1)
	go writerFunc(1, 100, mu, &wg)

	wg.Add(1)
	go writerFunc(2, 200, mu, &wg)

	// 再启动几个读取 goroutine，演示写操作完成后的读取
	time.Sleep(300 * time.Millisecond)
	for i := 5; i < 8; i++ {
		wg.Add(1)
		go readerFunc(i, mu, &wg)
	}

	// 等待所有 goroutine 执行完毕
	wg.Wait()
	printWithTime("Final data value: %d", sharedData)
	printWithTime("All goroutines completed")
}

// 写一个能体现 Once() 特性的测试函数
// 使用 sync.Once 来确保某个函数只执行一次
func testOnce() {
	var once sync.Once
	var count int

	increment := func() {
		count++
		printWithTime("Incremented count to %d", count)
	}

	// 多次调用 once.Do，但 increment 函数只会执行一次
	for i := 0; i < 5; i++ {
		once.Do(increment)
		printWithTime("once.Do called %d times", i+1)
	}

	printWithTime("Final count value: %d", count)
}

func testWaitGroup() {
	wg := &sync.WaitGroup{}
	var count int

	increment := func() {
		count++
		printWithTime("Incremented count to %d", count)
		wg.Done()
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go increment()
	}

	wg.Wait()
	printWithTime("Final count value: %d", count)
}

// 测试 map 的并发安全
// 本方法会报错 fatal error: concurrent map read and map write
func testMapPanic(skip bool) {
	if skip {
		return
	}
	m := make(map[string]int)
	go func() {
		for {
			m["key"] = 1
		}
	}()
	go func() {
		for {
			fmt.Println(m["key"])
		}
	}()
	time.Sleep(5 * time.Second)
}

func testMap() {
	mu := &sync.Mutex{}
	m := make(map[string]int)
	wg := &sync.WaitGroup{}

	increment := func(key string) {
		mu.Lock()
		defer mu.Unlock()
		m[key]++
		printWithTime("Incremented %s to %d", key, m[key])
		wg.Done()
	}

	for i := 0; i < 5; i++ {
		wg.Add(2)
		go increment(fmt.Sprintf("key%d", i))
		go increment("key")
	}

	wg.Wait()
	printWithTime("Final map value: %v", m)
}

func testSyncMap() {
	var m sync.Map
	wg := &sync.WaitGroup{}

	increment := func(key string) {
		defer wg.Done()
		actual, _ := m.LoadOrStore(key, 0)
		val := actual.(int) + 1
		m.Store(key, val)

		printWithTime("Incremented %s to %d", key, val)
	}

	for i := 0; i < 5; i++ {
		wg.Add(2)
		go increment(fmt.Sprintf("key%d", i))
		go increment("key")
	}

	wg.Wait()

	// 打印map
	m.Range(func(key, value any) bool {
		printWithTime("Final map value: key = %s, value = %d", key, value)
		return true
	})
}

type User struct {
	Name string
	Age  int
	Sex  bool
}

type Student struct {
	User
	Score int
}

func main() {
	startTime = time.Now()

	printWithTime("\n====== testMapPanic =======")
	testMapPanic(true) // 如果想看报错，请用false参数

	printWithTime("\n====== testMap =======")
	testMap()

	printWithTime("\n====== testSyncMap =======")
	testSyncMap()

	printWithTime("\n====== testOnce =======")
	testOnce()

	printWithTime("\n====== testWaitGroup =======")
	testWaitGroup()

	printWithTime("\n====== testMutex =======")
	testMutex()

	printWithTime("\n====== testRWMutex =======")
	testRWMutex()

}
