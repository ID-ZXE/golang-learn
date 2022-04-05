package main

import (
	"fmt"
)

//复杂处理的初始化，可以通过将初始化逻辑包装为一个匿名函数处理
var data [256]byte = func() (pc [256]byte) {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	return
}()

func init() {
	for _, item := range data {
		print(item, " ")
	}
	println()
}

func main() {
	// 修改外部值
	n := 0
	reply := &n
	Multiply(10, 5, reply)
	fmt.Println("Multiply:", *reply)
	// 多返回值
	var min, max int
	min, max = MinMax(78, 65)
	fmt.Printf("Minmium is: %d, Maximum is: %d\n", min, max)
}

// Multiply 修改外部值
func Multiply(a, b int, reply *int) {
	*reply = a * b
}

// MinMax 多个返回值
func MinMax(a int, b int) (min int, max int) {
	if a < b {
		min = a
		max = b
	} else {
		min = b
		max = a
	}
	return
}
