package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	// error 错误处理
	res, err := processError(100)
	if err == nil {
		println("res", res)
	} else {
		log.Println(err)
	}

	// defer 延迟执行
	// deferFunction1()
	// deferFunctionA()

	// recover
	defer func() {
		switch p := recover(); p {
		case nil: // no panic 无异常
		case "expected": // "expected" panic 可预期异常
			err = fmt.Errorf("expected panic")
		default:
			err = fmt.Errorf("unexpected panic")
		}
		// 格式化打印error
		fmt.Printf("%v\n", err)
	}()
	// panic
	// Go的类型系统会在编译时捕获很多错误，但有些错误只能在运行时检查，如数组访问越界、 空指针引用等。这些运行时错误会引起painc异常
	// 当panic异常发生时，程序会中断运行，并立即执行在该goroutine（可以先理解成 线程，在第8章会详细介绍）中被延迟的函数（defer 机制）。
	// 随后，程序崩溃并输出日志信 息。日志信息包括panic value和函数调用的堆栈跟踪信息
	occurPanic(3)
}

func processError(a int) (r int, err error) {
	if a > 10 {
		err = errors.New("this is error message")
		return
	}
	return a, nil
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

func occurPanic(x int) {
	// panics if x == 0
	fmt.Printf("f(%d)\n", x+0/x)
	defer fmt.Printf("defer %d\n", x)
	occurPanic(x - 1)
}
