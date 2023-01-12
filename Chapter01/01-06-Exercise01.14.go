// 從指標取得值

package main

import (
	"fmt"
	"time"
)

func main() {
	var count1 *int
	count2 := new(int)
	countTemp := 5
	count3 := &countTemp
	t := &time.Time{}

	if count1 != nil {
		// 用 * 取得指標的值
		fmt.Printf("count1: %#v \n", *count1)
	}

	if count2 != nil {
		fmt.Printf("count2: %#v \n", *count2)
	}

	if count3 != nil {
		fmt.Printf("count3: %#v \n", *count3)
	}

	if t != nil {
		fmt.Printf("time: %#v \n", *t)
		// 存取 t (time.Time 結構)
		// 自身的函式時不需要寫成 *t 來解除參照
		fmt.Printf("time: %#v \n", t.String())
	}
}