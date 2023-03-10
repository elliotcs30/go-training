// 使用 switch 敘述

package main

import (
	"fmt"
	"time"
)

func main() {
	// 定義變數並設定為星期的某一日
	// day := time.Sunday
	// day := time.Weekday(6)
	day := time.Now().Weekday()

	// 比對變數是星期幾
	switch day {
	case time.Monday: 
		fmt.Println("星期一, 猴子穿新衣")
	case time.Tuesday:
		fmt.Println("星期二, 猴子肚子餓")
	case time.Wednesday:
		fmt.Println("星期三, 猴子去爬山")
	case time.Thursday:
		fmt.Println("星期四, 猴子去考試")
	case time.Friday:
		fmt.Println("星期五, 猴子去跳舞")
	case time.Saturday:
		fmt.Println("星期六, 猴子去斗六")
	case time.Sunday:
		fmt.Println("星期日, 猴子過生日")
	default:
		fmt.Println("日期不正確")
	}
}