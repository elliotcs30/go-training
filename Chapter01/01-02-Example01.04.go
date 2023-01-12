// 以短變數宣告建立多重變數:

package main

import (
	"fmt"
	"time"
)

func main() {
	// 在同一行宣告多重變數
	Debug, LogLevel, startUpTime := false, "info", time.Now()
	fmt.Println(Debug, LogLevel, startUpTime)
}