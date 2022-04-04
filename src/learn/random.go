package main

import (
	"math/rand"
)
import "time"

func main() {
	var times int64 = int64(time.Now().Nanosecond())
	rand.Seed(times)
	// 换行
	println()
	println(rand.Int())
}
