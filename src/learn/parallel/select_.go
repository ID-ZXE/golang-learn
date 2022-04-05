package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// countdown()
	// countdown2()
	select2()
}

func countdown() {
	var abortChan chan int = make(chan int)
	fmt.Println("Commencing countdown. Press return to abortChan.")

	go abort(abortChan)

	// time.After函数会立即返回一个channel，并起一个新的goroutine在经过特定的时间后向该channel发送一个独立的值
	// 第二个case是指从abort管道读取数据
	select {
	case <-time.After(5 * time.Second):
	case <-abortChan:
		fmt.Println("Launch aborted!")
		return
	}
	fmt.Println("Launch")
}

func countdown2() {
	var abortChan chan int = make(chan int)
	fmt.Println("Commencing countdown. Press return to abort.")
	tick := time.Tick(1 * time.Second)

	go abort(abortChan)

	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
			// Do nothing.
		case <-abortChan:
			fmt.Println("Launch aborted!")
			return
		}
	}
	fmt.Println("Launch")
}

func abort(abortChan chan int) {
	rand.Seed(int64(time.Now().Nanosecond()))
	randomTime := int64(rand.Int() % 15)
	fmt.Printf("after %ds abort\n", randomTime)
	time.Sleep(time.Duration(randomTime) * time.Second)
	abortChan <- 1
}

// 结果永远是 0 2 4 6 8
func select2() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}
