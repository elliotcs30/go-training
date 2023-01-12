// 延伸習題: 訊息錯誤
// 請修正下列程式碼

package main

import "fmt"

func main() {
	count := 5
	message := ""
	if count > 5 {
		message = "Greater that 5"
	} else {
		message = "Not greater than 5"
	}

	fmt.Println(message)
}