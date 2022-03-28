package main

import "log"

type Stu struct {
	id   int
	name string
}

/**
无法更改id值 因为go是值传递
*/
func (s Stu) updateId(val int) {
	s.id = val
}

func (s *Stu) updateIdPoint(val int) {
	s.id = val
}

func (s Stu) compareStuId(stu Stu) bool {
	return s.id > stu.id
}

type Job struct {
	Command string
	*log.Logger
}

func (job *Job) start() {
	job.Println("start...")
}

type ReadWriter interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
}

type ReadWriterImpl struct {
}

func (readWriter ReadWriterImpl) Read(buf []byte) (n int, err error) {
	return 1, nil
}

func (readWriter ReadWriterImpl) Write(buf []byte) (n int, err error) {
	return 1, nil
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
	// stu3 := &Stu{id: 3, name: "c"}

	result := stu.compareStuId(stu2)
	println("result", result)

	println(stu2.id)
	stu2.updateIdPoint(4)
	println(stu2.id)
}

/**
实现了接口的全部方法 就是实现了接口
*/
func method2() {
	var obj ReadWriter = ReadWriterImpl{}
	println(obj)
}
