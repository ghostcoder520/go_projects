package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func test(num int) {
	defer wg.Done()
	for i := 1; i <= 3; i++ {
		fmt.Printf("协程%v打印了第%v条数据.\n", num, i)
		time.Sleep(time.Millisecond * 1000)
	}
}

func test2() {
	for i := 1; i < 3; i++ {
		fmt.Println("test2() 你好golang-", i)
		time.Sleep(time.Millisecond * 50)
	}

	wg.Done() // 协程计数器-1
}

var wg sync.WaitGroup

func demo1() {

	startTime := time.Now().Unix()
	//获取当前计算机CPU核心数
	fmt.Printf("当前计算机上的CPU个数: %d\n", runtime.NumCPU())
	//可以手动设定当前程序使用的CPU个数：
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	fmt.Printf("ok\n")

	for i := 1; i <= 5; i++ {
		wg.Add(1)  // 协程计数器+1
		go test(i) // 开启一个协程
	}
	wg.Wait()
	endTime := time.Now().Unix()

	fmt.Println("主线程执行完毕. 耗时:", endTime-startTime)
}

func demo2() {

	//1. 创建channel
	ch := make(chan int, 3) // int类型，容量为3的管道

	// 管道的长度和容量
	fmt.Printf("管道变量的值:%v 长度:%v 容量:%v\n", ch, len(ch), cap(ch)) //管道变量的值:0xc00006a100 长度:0 容量:3
	fmt.Printf("%T\n", ch)                                      // chan int

	//2. 往管道里面存数据
	ch <- 10 // 长度1
	ch <- 20 // 长度2
	ch <- 30 // 长度3

	//3. 取出管道内的数据

	a := <-ch
	fmt.Println(a) //10

	<-ch //20

	c := <-ch
	fmt.Println(c) //30

	//4. 管道阻塞: 管道放满时继续放数据，或者管道为空时取数据（像一个阻塞队列）

	ch1 := make(chan int, 1)

	ch1 <- 10
	// ch1 <- 20 // all goroutines are asleep - deadlock!
	<-ch1
	// <- ch1 // all goroutines are asleep - deadlock!

	// 循环遍历管道
	{
		ch1 := make(chan int, 10)
		for i := 1; i < 10; i++ {
			ch1 <- i
		}

		close(ch1) // 关闭管道

		// for range时 必须先关闭管道，否则死锁
		for v := range ch1 { // 管道只有value，没有key
			fmt.Printf("%v ", v)

		}
		fmt.Println()

		// 管道没关闭时只能通过 普通for来进行遍历

	}

}

var wg1 sync.WaitGroup

// 写数据
func fn1(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("【写入】管道:%v\n", i)
		time.Sleep(time.Millisecond * 50)
		ch <- i
	}
	close(ch)
	fmt.Println("管道已关闭.")
	wg1.Done()
}

// 读数据
func fn2(ch <-chan int) {
	for v := range ch {
		fmt.Printf("【读取】管道:%v\n", v)
		time.Sleep(time.Millisecond * 50)
	}
	wg1.Done()
}

func demo3() {
	ch := make(chan int, 10)

	wg1.Add(1)
	go fn1(ch)
	wg1.Add(1)
	go fn2(ch)

	wg1.Wait()

	fmt.Println("主线程结束.")
}

func demo4() {

	intChan := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		intChan <- i
	}

	strChan := make(chan string, 10)
	for i := 1; i <= 10; i++ {
		strChan <- "hello" + fmt.Sprintf("%d", i)
	}

	for {
		select {
		case v := <-intChan:
			fmt.Printf("从intChan读取到数据: %v\n", v)
			// time.Sleep(time.Millisecond * 50)
		case v := <-strChan:
			fmt.Printf("从strChan读取到数据: %v\n", v)
			// time.Sleep(time.Millisecond * 50)
		default:
			fmt.Println("数据获取完毕")
			return
		}
	}
}

var count = 0
var wg2 sync.WaitGroup
var mutex sync.Mutex

func myInc(num int) {

	mutex.Lock()

	count++
	fmt.Printf("[协程 %02d] the count is: %v\n", num, count)

	mutex.Unlock()

	wg2.Done()

}

func demo5() {
	for i := 1; i <= 20; i++ {
		wg2.Add(1)
		go myInc(i)
	}
	wg2.Wait()
}

var myLock sync.RWMutex
var wg3 sync.WaitGroup

func read(num int) {
	myLock.RLock()
	fmt.Printf("[读协程 %02d]\n", num)
	time.Sleep(time.Millisecond * 1000)
	myLock.RUnlock()

	wg3.Done()
}

func write(num int) {
	myLock.Lock()
	fmt.Printf("[写协程 %02d]\n", num)
	time.Sleep(time.Millisecond * 1000)
	myLock.Unlock()

	wg3.Done()
}

func demo6() {

	for i := 1; i <= 10; i++ {
		wg3.Add(1)
		go read(i)
	}

	for i := 1; i <= 10; i++ {
		wg3.Add(1)
		go write(i)
	}

	wg3.Wait()

}

func main() {

	// 每个demo独立运行：

	// demo1()  //goroutine概念
	// demo2()  // 管道概念
	// demo3()
	// demo4() // select
	// demo5() // 互斥锁
	demo6()

}
