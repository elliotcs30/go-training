// 延伸習題：計數錯誤
// 請修正下列程式碼

package main

import "fmt"

func main() {
	count := 0
	if count < 5 {
		count = 10
		count++
	}
	fmt.Println(count == 11)
}