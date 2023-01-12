package main

import (
	"fmt"
	"time"
)

// 一次宣告多個變數
var (
	Debug bool = false
	LogLevel string = "info"
	startUpTime time.Time = time.Now()
)

func main() {
	fmt.Println(Debug, LogLevel, startUpTime)
}