package main

import (
	"fmt"
	"os"
	"unicode"
)

// PI 常量
const PI float32 = 3.1415

// 常量自动递增
const (
	Unknown = iota
	Female
	Male
)

// int变量
var i int = 100

// 字符串变量
var str string = "hello world"

var (
	HOME   = os.Getenv("HOME")
	USER   = os.Getenv("USER")
	GOROOT = os.Getenv("GOROOT")
)

// 在main函数之前执行
func init() {
	println()
	println("home", HOME, "user", USER, "go root", GOROOT)
}

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
