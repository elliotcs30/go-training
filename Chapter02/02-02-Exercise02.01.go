// if 敘述

package main

import "fmt"

func main() {
	// 定義一個型別為 int 的變數, 並賦予起始值
	// 將其指定為 5, 它是奇數, 也可以換成偶數
	input := 5

	// 以 if 敘述檢查除以 2 的餘數是否為 0
	if input%2 == 0 {
		fmt.Println(input, "是偶數")
	}
	if input%2 == 1 {
		fmt.Println(input, "是奇數")
	}
}