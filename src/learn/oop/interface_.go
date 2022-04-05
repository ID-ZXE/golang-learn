package main

import (
	"bytes"
	"io"
)

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

/*
ReadWriter 组合Reader与Writer接口
我们可以用这种方式以一个简写命名另一个接口，而不用声明它所有的方法。这种方式本称为接口内嵌
*/
type ReadWriter interface {
	Reader
	Writer
}

// ReadWriteCloser 组合ReadWriter与Closer接口
type ReadWriteCloser interface {
	ReadWriter
	Closer
}

func main() {
	impl()
	// testNPE()
}

/************************/

type ReadWriterImpl struct {
}

func (readWriter ReadWriterImpl) Read(buf []byte) (n int, err error) {
	return 1, nil
}

func (readWriter ReadWriterImpl) Write(buf []byte) (n int, err error) {
	return 1, nil
}

/**
一个类型如果拥有一个接口需要的所有方法，那么这个类型就实现了这个接口
*/
func impl() {
	var obj ReadWriter = ReadWriterImpl{}
	println(&obj)
	// 类型断言
	writerImpl := obj.(ReadWriterImpl)
	println(&writerImpl)
}

/************************/

func testNPE() {
	var buf *bytes.Buffer
	// true
	println(buf == nil)
	// (out != nil) == true
	// 是out变量是一个包含空指针值的非 空接口
	f(buf)

	var buf2 io.Writer
	// true
	println(buf == nil)
	// (out != nil) == false
	f(buf2)
}

func f(out io.Writer) {
	// ...do something...
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}
