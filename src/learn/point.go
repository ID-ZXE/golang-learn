package main

import "fmt"

func main() {
	var i1 int = 5
	fmt.Printf("An integer: %d, it’s location in memory: %p\n", i1, &i1)

	// 取地址符号
	var intP *int = &i1
	println("address", intP)

	s := "good bye"
	var p *string = &s
	*p = "ok!!!"
	// prints address
	fmt.Printf("Here is the pointer p: %p\n", p)
	// prints string
	fmt.Printf("Here is the string *p: %s\n", *p)
	// prints same string
	fmt.Printf("Here is the string s: %s\n", s)

	// 调用f函数时创建局 部变量v，在局部变量地址被返回之后依然有效，因为指针p依然引用这个变量
	var p2 = f()
	println(p2)
	println(f() == f())

	v := 1
	incr(&v)
	// side effect: v is now 2
	fmt.Println(incr(&v)) // "3" (and v is 3)
}

func f() *int {
	v := 1
	return &v
}

// ++i操作
func incr(p *int) int {
	// 非常重要：只是增加p指向的变量的值，并不改变p指针！！！
	*p++
	return *p
}
