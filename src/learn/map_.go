package main

import "fmt"

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {
	operate()
	// mockKvDB()
}

func operate() {
	// 三种定义方式
	// ages := map[string]int{}
	// ages := make(map[string]int)
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	var name = "alice"
	fmt.Printf("%s age is %d\n", name, ages[name])

	for _name, _age := range ages {
		fmt.Printf("_name:%s, _age:%d\n", _name, _age)
	}

	delete(ages, name)
	if val, ok := ages[name]; !ok {
		// delete之后val为0
		fmt.Printf("after delete. %s age val is %d", name, val)
	}
}

func mockKvDB() {
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
