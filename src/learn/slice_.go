package main

func init() {
	println()
}

// Slice（切片）代表变长的序列，序列中每个元素都有相同的类型
// 一个slice类型一般写作[]T，其中T代表slice中元素的类型；slice的语法和数组很像，只是没有固定长度而已。
func main() {
	// 定义数组
	// arr := [3]int{1, 2, 3}
	// 定义切片
	// slice := []int{1, 2, 3}
	// arr = slice
	// compile error: Cannot use 'slice' (type []int) as the type [3]int

	// slice1()
	// slice2()
	// slice3()
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

// slice唯一合法的比较操作是和nil比较
func compare() {
	var s []int    // len(s) == 0, s == nil
	s = nil        // len(s) == 0, s == nil
	s = []int(nil) // len(s) == 0, s == nil
	s = []int{}    // len(s) == 0, s != nil

	println(s == nil)
}

func stack() {
	stack := make([]int, 0)
	// push v
	stack = append(stack, 1)
	// top of stack
	// top := stack[len(stack)-1]
	// pop
	stack = stack[:len(stack)-1]
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
