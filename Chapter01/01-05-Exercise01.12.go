// 零值(zero values)
// 印出零值:

package main

import (
	"fmt"
	"time"
)

func main() {
	// 整數, count 的值會帶入前面字串的 %#v
	var count int
	fmt.Printf("Count : %#v \n", count)

	// 64 位元浮點數
	var discount float64
	fmt.Printf("Discount : %#v \n", discount)

	// 布林值
	var debug bool
	fmt.Printf("Debug : %#v \n", debug)

	// 字串
	var message string
	fmt.Printf("Message : %#v \n", message)

	// 字串切片
	var emails []string
	fmt.Printf("Emails : %#v \n", emails)

	// time.Time 結構
	var startTime time.Time
	fmt.Printf("Start : %#v \n", startTime)
}