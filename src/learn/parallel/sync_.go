package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mu      sync.Mutex
	balance int
)

var (
	mu2      sync.RWMutex
	balance2 int
)

// sync.Once

func main() {
	var x, y int

	go func() {
		// A1
		x = 1
		// A2
		fmt.Print("y:", y, " ")
	}()

	go func() {
		// B1
		y = 1
		// B2
		fmt.Print("x:", x, " ")
	}()

	time.Sleep(1 * time.Second)
}

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	balance = balance + amount
}

func Balance() int {
	mu.Lock()
	defer mu.Unlock()
	return balance
}

// go中没有可重入锁 下面的操作会发生死锁
func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	Deposit(-amount)
	if Balance() < 0 {
		Deposit(amount)
		return false
	}
	return true
}

func Deposit2(amount int) {
	mu2.Lock()
	defer mu2.Unlock()
	balance2 = balance2 + amount
}

func Balance2() int {
	mu2.RLock()
	defer mu2.RUnlock()
	return balance2
}
