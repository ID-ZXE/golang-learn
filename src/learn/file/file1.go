package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
)

var path string = os.Getenv("HOME") + "/" + "/file/tmp/"

func main() {
	// 追加模式
	// os.O_APPEND
	// 读写模式
	// os.O_RDWR

	filePath := path + "test1.txt"
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0777)
	defer file.Close()
	if err != nil {
		fmt.Printf("OpenFile err:%v\n", err)
	}

	write(file)

	//read(file)
	//seek2ZeroOffset(file)
	//read(file)
	//seek2ZeroOffset(file)

	//bufioOpe(file)

	//data, _ := os.ReadFile(filePath)
	//fmt.Println("data:", string(data))

	//dirOpe()
}

func read(file *os.File) {
	byteArr := make([]byte, 5)
	for {
		readLength, err := file.Read(byteArr)
		if err != nil {
			fmt.Printf("\nerr:%v\n", err)
			return
		}
		if readLength > 0 {
			fmt.Print(string(byteArr[:readLength]))
		}
	}
}

func seek2ZeroOffset(file *os.File) {
	file.Seek(0, io.SeekStart)
}

func write(file *os.File) {
	kv := make(map[string]string)
	kv["name"] = "test"
	data, err := json.Marshal(kv)
	if err != nil {
		fmt.Printf("OpenFile err:%v\n", err)
	}
	file.Write(data)
}

func bufioOpe(file *os.File) {
	reader := bufio.NewReader(file)
	writer := bufio.NewWriter(file)
	n := 0
	for {
		n++
		// isPrefix 当前行是否读完
		line, isPrefix, err := reader.ReadLine()
		if err != nil {
			fmt.Printf("err:%v\n", err)
			break
		}
		fmt.Println(string(line), isPrefix)
		// 写入缓存 并不写入硬盘
		writer.WriteString(strconv.Itoa(n) + " " + string(line))
	}
	// 刷新到缓存
	seek2ZeroOffset(file)
	writer.Flush()
}

func dirOpe() {
	dirs, _ := os.ReadDir("./")
	for _, dir := range dirs {
		info, _ := dir.Info()
		fmt.Println(info.Name(), info.IsDir())
	}
}
