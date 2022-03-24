package main

import (
	"fmt"
	"os"
	"unicode"
)

// PI 显式的定义常量
const PI float32 = 3.1415

// STR1 隐式类型定义常量
const STR1 = "str1"

// 常量自动递增
const (
	Unknown = iota
	Female
	Male
)

const (
	Monday, Tuesday, Wednesday = 1, 2, 3
	Thursday, Friday, Saturday = 4, 5, 6
)

// i int变量
var i int = 100

// str 字符串变量
var str string = "hello world"

var (
	HOME   = os.Getenv("HOME")
	USER   = os.Getenv("USER")
	GOROOT = os.Getenv("GOROOT")
)

/**
在每个包初始化之后执行，不能被人为调用
*/
func init() {
	println()
	println("home", HOME, "user", USER, "go root", GOROOT)
}

/**
程序入口
*/
func main() {
	println(str)
	method1()
	char()
	strAppend()
}

func method1() {
	// 这个变量如果没有使用 会报错
	var a, b int
	a, b = method2(1, 2)
	println(a, b)

	// 声明并赋值, 无需var
	// 这个变量如果没有使用 会报错
	num := 10
	println("num", num)

	i, j := 10, 11
	println(i, j)
}

func method2(param1 int, param2 int) (result1 int, result2 int) {
	println(param1, param2)
	return 100, 1000
}

func char() {
	var char1 int = 66
	fmt.Printf("char %c\n", char1)

	// 判断是否为字母
	println(unicode.IsLetter(rune(char1)))
	// 判断是否为数字
	println(unicode.IsDigit(rune(char1)))
	// 判断是否为空白符号
	println(unicode.IsSpace(rune(char1)))
}

// 类型转换
func cast() {
	var a int16 = 32
	var b int32 = 64
	a = int16(b)
	println(a, b)
}

// 字符串拼接
func strAppend() {
	str := "s" + "t" + "r"
	println(str)
}
