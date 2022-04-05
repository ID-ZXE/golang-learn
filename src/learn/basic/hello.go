package main

import (
	"os"
)

// PI 显式的定义常量
const PI float32 = 3.1415

// STR1 隐式类型定义常量
const STR1 = "str1"

// 常量利用iota自动递增
const (
	Unknown = iota
	Female
	Male
)

// iota被重置了
const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	GiB // 1073741824
	TiB // 1099511627776 (exceeds 1 << 32)
	PiB // 1125899906842624
	EiB // 1152921504606846976
	ZiB // 1180591620717411303424 (exceeds 1 << 64)
	YiB // 1208925819614629174706176
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
	println(str, MiB)
	valTest()
	appendTest()
	arrTest()
}

func valTest() {
	// 这个变量如果没有使用 会报错
	var a, b int
	a, b = method(1, 2)
	println(a, b)

	// 声明并赋值, 无需var
	// 这个变量如果没有使用 会报错
	num := 10
	println("num", num)

	i, j := 10, 11
	println(i, j)
}

func method(param1 int, param2 int) (result1 int, result2 int) {
	println(param1, param2)
	return 100, 1000
}

// 类型转换
func cast() {
	var a int16 = 32
	var b int32 = 64
	a = int16(b)
	println(a, b)
}

// 字符串拼接
func appendTest() {
	println("s" + "t" + "r")
}

func arrTest() {
	// 默认情况下，数组的每个元素都被初始化为元素类型对应的零值，对于数字类型来说就是0
	var arr [16]int
	arr[15] = 100
	println(arr[0], arr[15])

	a := [2]int{1, 2}
	b := [...]int{1, 2}
	// 如果一个数组的元素类型是可以相互比较的，那么数组类型也是可以相互比较的
	// 这时候我们可以直接通过==比较运算符来比较两个数组，只有当两个数组的所有元素都是相等的时候 数组才是相等的
	// result=true
	println(a == b)

	// c := [3]int{1, 2}
	// println(a == c)
	// compile error: cannot compare [2]int == [3]int
}
