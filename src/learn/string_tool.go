package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func init() {
	println()
}

func main() {
	// stringsTool()
	strconvTool()
	// unicodeTool()
	transfer()
	println(intsToString([]int{1, 2, 3}))
}

// strings包 提供了许多如字符串的查询、替换、比较、截断、拆分和合并等功能
func stringsTool() {
	var str string = "This is an example of string"
	println("HasPrefix", strings.HasPrefix(str, "Th"))
	println("str[0]", str[0])
	println("len", len(str))

	println("isContain:", strings.Contains(str, "example"))
	println("index:", strings.Index(str, "example"))
	println("lastIndex:", strings.LastIndex(str, "example"))
	// 如果 ch 是非 ASCII 编码的字符，建议使用以下函数来对字符进行定位
	println("index:", strings.IndexRune(str, rune('B')))

	println("repeat s result:", strings.Repeat("s", 3))
	// -1表示替换所有 不改变原有字符串
	println("replace result:", strings.Replace(str, "This", "this", -1))
	println("example count:", strings.Count(str, "example"))
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

// strconv包提供了布尔型、整型数、浮点数和对应字符串的相互转换，还提供了双引号转义相关的转换
func strconvTool() {
	// 返回数字i所表示的字符串类型的十进制数
	num := 123
	var str1 string = strconv.Itoa(num)
	var str2 string = fmt.Sprintf("%d", num)
	println(str1, str2)

	// 将字符串转为int
	// val, err := strconv.Atoi("1")
	var val int
	val, _ = strconv.Atoi(str1)
	println(val)
}

// unicode包提供了IsDigit、IsLetter、IsUpper和IsLower等类似功能，它们用于给字符分类。
// 每个函数有一个单一的rune类型的参数，然后返回一个布尔值。
// 而像ToUpper和ToLower之类的转换函数将用于rune字符的大小写转换。
func unicodeTool() {
	var c int = 66
	fmt.Printf("char %c\n", c)

	// 判断是否为字母
	println(unicode.IsLetter(rune(c)))
	// 判断是否为数字
	println(unicode.IsDigit(rune(c)))
	// 判断是否为空白符号
	println(unicode.IsSpace(rune(c)))
}

func transfer() {
	s := "abc"
	// 从概念上讲，一个[]byte(s)转换是分配了一个新的字节数组用于保存字符串数据的拷贝
	b := []byte(s)
	// 但总的来说需要确保在变量b被修改的情况下，原始的s字符串也不会改变。
	// 将一个字节slice转到字符串的string(b)操作则是构造一个字符串拷贝，以确保s2字符串是只读的。
	s2 := string(b)
	println(s2)
}

// bytes包还提供了Buffer类型用于字节slice的缓存。一个Buffer开始是空的
// 但是随着string、 byte或[]byte等类型数据的写入可以动态增长
// 一个bytes.Buffer变量并不需要初始化，因为零值也是有效的
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
