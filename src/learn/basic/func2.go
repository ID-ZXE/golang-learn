package main

import "fmt"

func main() {
	// 可变长参数
	println(sum(1, 2, 3))
	// 匿名函数
	f := func(x, y int) int {
		return x + y
	}
	println(f(1, 2))
	// 函数参数
	f2 := func(i int) int {
		return i * i
	}
	println(funcArg(f2))
	// 闭包
	callSquares()
}

// sum 可变长参数
func sum(ints ...int) int {
	result := 0
	for _, val := range ints {
		result += val
	}
	return result
}

// 函数也可以作为参数
func funcArg(f func(n int) int) int {
	return f(10)
}

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

// 函数值不仅仅是一串代码，还记录了状态。在squares中定义的匿名内 部函数可以访问和更新squares中的局部变量，
// 这意味着匿名函数和squares中，存在变量引用。
// 这就是函数值属于引用类型和函数值不可比较的原因。
// Go使用闭包（closures）技术实现函数值，Go程序员也把函数值叫做闭包。
func callSquares() {
	f := squares()
	fmt.Println(f()) // "1"
	fmt.Println(f()) // "4"
	fmt.Println(f()) // "9"
	fmt.Println(f()) // "16"
}
