package main

func init() {
	println()
}

func main() {
	// slice1()
	// slice2()
	slice3()
}

func slice1() {
	var str string = "hello,golang"
	// 切片到5位置
	println(str[:5])
	// 切片从第6个开始, 不包括第6个
	println(str[6:])
	// 切片从第4个开始, 不包括第4个, 切到第7个
	println(str[4:7])
	// 所有
	println(str[:])
}

func slice2() {
	var mySlice = make([]int, 5, 10)
	println("mySlice[4]", mySlice[4])
	println("len(mySlice):", len(mySlice))
	println("cap(mySlice):", cap(mySlice))
	mySlice = append(mySlice, 1, 2, 3)
	println(mySlice[5], mySlice[6], mySlice[7])
	mySlice1 := []int{4, 5}
	mySlice = append(mySlice, mySlice1...)
	println(mySlice[8], mySlice[9])

	// copy mySlice前三个到mySlice1中
	copy(mySlice1, mySlice)
	println("mySlice1", mySlice1[0], mySlice1[1])
}

func slice3() {
	var arr = make([]int, 0)
	arr = append(arr, 1, 2, 3)
	for _, item := range arr {
		println(item)
	}
}
