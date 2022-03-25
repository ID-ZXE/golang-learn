package main

import "fmt"

func init() {
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

	// defer 延迟执行
	deferFunction1()
	deferFunctionA()

	multiArg("a", "b", "c")

	// 匿名函数
	f := func(x, y int) int {
		return x + y
	}
	println(f(1, 2))
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

// multiArg 可变长参数
func multiArg(str ...string) {
	println(str[0])
}

func deferFunction1() {
	fmt.Printf("In Function1 at the top\n")
	// 推迟function2的执行
	defer deferFunction2()
	fmt.Printf("In Function1 at the bottom!\n")
}
func deferFunction2() {
	fmt.Printf("Function2: Deferred until the end of the calling function!\n")
}

func deferFunctionA() {
	defer un(trace("functionA"))
	fmt.Println("in functionA")
	deferFunctionB()
}
func deferFunctionB() {
	defer un(trace("functionB"))
	fmt.Println("in functionB")
}
func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}
func un(s string) {
	fmt.Println("leaving:", s)
}
