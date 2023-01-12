// 透過函式傳回值宣告多重變數

package main
import (
	"fmt"
	"time"
)

func getConfig() (bool, string, time.Time) {
	return false, "info", time.Now()
}

func main() {
	// 用函式的傳回值作為初始值:
	Debug, LogLevel, startUpTime := getConfig()
	fmt.Println(Debug, LogLevel, startUpTime)

	// var 函式回傳值:
	// var (
	// 	Debug bool
	// 	LogLevel string
	// 	startUpTime time.Time
	// )

	// Debug, LogLevel, startUpTime = getConfig()
	// fmt.Println(Debug, LogLevel, startUpTime)
}