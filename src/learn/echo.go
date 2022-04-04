package main

import (
	"flag"
	"fmt"
	"strings"
)

// 获取命令行参数  seq 默认值为空格
var sep = flag.String("s", " ", "separator")

// 获取命令行参数 n
var n = flag.Bool("n", false, "omit trailing newline")

func main() {
	// Flag解析在第一个非flag参数（单个”-“不是flag参数）之前停止，或者在终止符”–“之后停止。
	flag.Parse()
	// -n -s - a b c
	// print a-b-c
	fmt.Print(strings.Join(flag.Args(), *sep))

	if !*n {
		fmt.Println()
	}

	fmt.Println("\nend")
}
