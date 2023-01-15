// nil 值

package main

import "fmt"

func main() {
	// 沒有初始值的指標數會是 nil
	var message *string

	if message == nil {
		fmt.Println("錯誤, 非預期的 nil 值")
		return
	}
	fmt.Println(&message)
}