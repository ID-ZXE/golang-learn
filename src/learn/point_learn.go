package main

import "fmt"

func main() {
	var i1 int = 5
	fmt.Printf("An integer: %d, it’s location in memory: %p\n", i1, &i1)

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
}
