// rune 型別

package main

import "fmt"

func main() {
	username := "Sir_King_Über"
	runes := []rune(username) // 將字串轉成 rune 切片

	for i := 0; i < len(runes); i++ {
		fmt.Print(string(runes[i]), " ") // 將 rune 轉為字串印出
	}
}