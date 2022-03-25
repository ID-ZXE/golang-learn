package main

import "fmt"

func main() {
	for_()
}

func for_() {
	var str string
	str = "hello, world"
	for i := 0; i < len(str); i++ {
		// 打印byte
		print(str[i], " ")
	}
	println()
	for i := 0; i < len(str); i++ {
		// 打印byte
		fmt.Printf("%c", str[i])
	}
	println()
	for _, v := range str {
		fmt.Printf("%c", v)
	}
}

func if_() {
	var first int = 10
	var cond int

	if first <= 0 {
		fmt.Printf("first is less than or equal to 0\n")
	} else if first > 0 && first < 5 {
		fmt.Printf("first is between 0 and 5\n")
	} else {
		fmt.Printf("first is 5 or greater\n")
	}

	if cond = 5; cond > 10 {
		fmt.Printf("cond is greater than 10\n")
	} else {
		fmt.Printf("cond is not greater than 10\n")
	}
}

func switch_() {
	var num1 int = 100

	switch num1 {
	case 98, 99:
		fmt.Println("It's equal to 98")
	case 100:
		fmt.Println("It's equal to 100")
	default:
		fmt.Println("It's not equal to 98 or 100")
	}
}
