package main

import "fmt"

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {
	mapProcess()
}

func mapProcess() {
	var personDB map[string]PersonInfo
	personDB = make(map[string]PersonInfo)

	// 往这个map里插入几条数据
	personDB["12345"] = PersonInfo{"12345", "Tom", "Room 203,..."}
	personDB["1"] = PersonInfo{"1", "Jack", "Room 101,..."}

	// 从这个map查找键为"1234"的信息
	person, ok := personDB["1234"]
	if ok {
		fmt.Println("Found person", person.Name, "with ID 1234.")
	} else {
		fmt.Println("Did not find person with ID 1234.")
	}
	// 删除
	delete(personDB, "1")
}
