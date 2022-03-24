package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "This is an example of functionB functionB string"
	println()
	fmt.Printf("T/F? Does the string \"%s\" have prefix %s? ", str, "Th")
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
	println()

	println("isContain:", strings.Contains(str, "example"))
	println("index:", strings.Index(str, "example"))
	println("lastIndex:", strings.LastIndex(str, "example"))
	// 如果 ch 是非 ASCII 编码的字符，建议使用以下函数来对字符进行定位
	println("index:", strings.IndexRune(str, rune('B')))

	println("repeat s result:", strings.Repeat("s", 3))
	// -1表示替换所有 不改变原有字符串
	println("replace result:", strings.Replace(str, "This", "this", -1))
	println("functionB count:", strings.Count(str, "functionB"))
	println("lower", strings.ToLower("ABC"), "upper", strings.ToUpper("abc"))

	println("****修剪字符串****")
	println("trim result", strings.TrimSpace(" abc"))
	println("trim result", strings.Trim("abc", "c"))

	println("****切割字符串****")
	var arr []string = strings.Split("abc-ijk", "-")
	println("index0", arr[0], "index1", arr[1])

	println("****拼接字符串****")
	join := strings.Join(arr, "+")
	println("join result", join)

	// 从字符串读取内容
	// reader := strings.NewReader(str)
}
