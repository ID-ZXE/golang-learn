package main

import (
	"strconv"
)

func main() {
	// 返回数字i所表示的字符串类型的十进制数
	var strs string = strconv.Itoa(67)
	println(strs)

	// 将字符串转为int
	// val, err := strconv.Atoi("1")
	var val int
	val, _ = strconv.Atoi("1")
	println(val)
}
