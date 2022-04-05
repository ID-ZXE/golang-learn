package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 一个channel有发送和接受两个主要操作，都是通信行为。一个发送语句将一个值从一个 goroutine通过channel发送到另一个执行接收操作的goroutine。
// 发送和接收两个操作都使用 <- 运算符。在发送语句中，
// <- 运算符分割channel和要发送的值。在接收语句中，
// <- 运算符写在channel对象之前。一个不使用接收结果的接收操作也是合法的

// ch <- x   // a send statement
// x = <-ch  // a receive expression in an assignment statement
// <-ch      // a receive statement; result is discarded

// Channel还支持close操作，用于关闭channel，随后对基于该channel的任何发送操作都将导 致panic异常。
// 对一个已经被close过的channel进行接收操作依然可以接受到之前已经成功发送的数据；如果channel中已经没有数据的话将产生一个零值的数据
// 使用内置的close函数就可以关闭一个channel
// close(ch)

// 一个基于无缓存Channels的发送操作将导致发送者goroutine阻塞，直到另一个goroutine在相同的Channels上执行接收操作，
// 当发送的值通过Channels成功传输之后，两个goroutine可以继续执行后面的语句。
// 反之，如果接收操作先发生，那么接收者goroutine也将阻塞，直到有另一个goroutine在相同的Channels上执行发送操作。
// ch = make(chan int)    // unbuffered channel
// ch = make(chan int, 0) // unbuffered channel

// 向缓存Channel的发送操作就是向内部缓存队列的尾部插入元素，接收操作则是从队列的头部删除元素。
// 如果内部缓存队列是满的，那么发送操作将阻塞直到因另一个goroutine执行接收 操作而释放了新的队列空间。
// 相反，如果channel是空的，接收操作将阻塞直到有另一个 goroutine执行发送操作而向队列插入元素。
// ch = make(chan int, 3) // buffered channel with capacity 3

func main() {
	// pipeline1()
	// pipeline2()
	println(mirroredQuery())
}

func pipeline1() {
	naturals := make(chan int)
	squares := make(chan int)

	num := 100

	// Counter
	go func() {
		for x := 0; x < num; x++ {
			naturals <- x
		}
		println("close naturals")
		close(naturals)
	}()

	// Squarer
	go func() {
		for {
			x, ok := <-naturals
			// channel was closed and drained
			if !ok {
				break
			}
			squares <- x * x
		}
		println("close squares")
		close(squares)
	}()

	// Printer (in main goroutine)
	for i := 0; i < num; i++ {
		println(i, <-squares)
	}
	// 主函数退出时 所有goroutine都会退出
	time.Sleep(100)
	println("end")
}

/****************/

func pipeline2() {
	naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}

// 单方向channel
func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

/****************/

// 获取最快的响应
func mirroredQuery() string {
	// 如果我们使用了无缓存的channel，那么两个慢的goroutines将会因为没有人接收而被永远卡住。
	// 这种情况，称为goroutines泄漏，这将是一个BUG。
	// 和垃圾变量不同，泄漏的goroutines 并不会被自动回收，因此确保每个不再需要的goroutine能正常退出是重要的。
	responses := make(chan string, 3)
	go func() { responses <- request("asia.gopl.io") }()
	go func() { responses <- request("europe.gopl.io") }()
	go func() { responses <- request("americas.gopl.io") }()
	// return the quickest response
	return <-responses
}

func request(hostname string) (response string) {
	sleep()
	return "result:" + hostname
}

func sleep() {
	rand.Seed(int64(time.Now().Nanosecond()))
	randomTime := int64(rand.Int() % 20)
	println("sleep", randomTime, "ms")
	// nanosecond
	time.Sleep(time.Duration(randomTime) * time.Microsecond)
}
