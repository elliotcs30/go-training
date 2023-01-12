// 取得指標:

package main

import (
	"fmt"
	"time"
)

func main() {
	// 宣告指標變數 (值為 nil)
	var count1 *int

	// 宣告指標變數 (值為 0)
	count2 := new(int)
	
	countTemp := 5

	// 從別的變數建立指標
	count3 := &countTemp

	// 直接從結構型別建立指標
	t := &time.Time{}

	fmt.Printf("count1: %#v \n", count1)
	fmt.Printf("count2: %#v \n", count2)
	fmt.Printf("count3: %#v \n", count3)
	fmt.Printf("time: %#v \n", t)

}