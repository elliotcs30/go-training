// 用其他變數來賦值:

package main

import "fmt"

var defaultOffset = 10

func main() {
	offset := defaultOffset
	fmt.Println(offset)

	// 把 offset 的值加上 defaultOffset 的值, 重新存入 offset
	offset = offset + defaultOffset
	fmt.Println(offset)
}