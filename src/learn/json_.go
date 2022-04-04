package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Movie
// “omitempty”选项指定如果字段具有空值，定义为false，0，零指针，nil接口值以及任何空数组，切片，映射或字符串，则该字段应从编码中省略。
type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

func main() {
	data, err := json.Marshal(movies)
	if err != nil {
		// Fatal is equivalent to Print() followed by a call to os.Exit(1).
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	// 匿名struct
	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
}
