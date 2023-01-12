// 一次更改多個變數值
package main

import "fmt"

func main() {
	// 宣告多重變數並賦予初始值
	query, limit, offset := "bat", 10, 0

	// 以單行敘述一次更改全部變數的值
	query, limit, offset = "ball", offset, 20

	fmt.Println(query, limit, offset)
}