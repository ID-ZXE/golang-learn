package main

import "log"

// Stu 如果结构体成员名字是以大写字母开头的，那么该成员就是public的
type Stu struct {
	id          int
	name        string
	Description string
}

/**
无法更改id值 因为go是值传递
*/
func (s Stu) updateId(val int) {
	s.id = val
}

/**
直接通过地址更改id值
如果要在函数内部修改结构体成员的话，用指针传入是必须的；
因为在Go语言中，所有的函数参数都是值拷贝传入的，函数参数将不再是函数调用时的原始变量。
*/
func (s *Stu) updateIdPoint(val int) {
	s.id = val
}

func (s *Stu) nilValue() int {
	if s == nil {
		return 0
	}
	return s.id
}

func (s Stu) compareStuId(stu Stu) bool {
	return s.id > stu.id
}

// Job *log.Logger 相当于拥有了log.Logger的所有属性与方法
// 匿名成员 可以直接访问叶子属性而不需要给出完整的路径
type Job struct {
	Command string
	*log.Logger
}

func (job *Job) start() {
	job.Println("start...")
}

func init() {
	println()
}

func main() {
	method()

}

func method() {
	var stu Stu = Stu{id: 1, name: "a"}
	var stu2 Stu = Stu{id: 2, name: "b"}

	// new关键字返回的是指针
	stuPoint := new(Stu)
	stuPoint.name = "point"
	println(stuPoint)

	result := stu.compareStuId(stu2)
	println("result", result)

	// 显式的传入指针
	r := &Stu{3, "c", "des"}
	r.updateIdPoint(1)

	// 隐式的类型转换 因为接收器是指针类型
	println(stu2.id)
	// 相当于 (&stu2).updateIdPoint(4)
	stu2.updateIdPoint(4)
	println(stu2.id)

	// nilValue可以接受nil值
	stu3 := &Stu{id: 4, name: "d"}
	// stu2不可以直接赋值nil 因为不是指针
	// 需要进行类型转换 (&stu2).nilValue()
	stu3 = nil
	println("nil val", stu3.nilValue())
}
