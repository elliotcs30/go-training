// 省略 var 宣告中的型別或賦值

package main

import (
	"fmt"
	"time"
)

var (
	Debug bool
	LogLevel = "info"
	startUpTime = time.Now()
)

func main() {
	fmt.Println(Debug, LogLevel, startUpTime)
}