// 檢查字串長度

package main

import "fmt"

func main() {
	username := "Sir_King_Über"

	// 取得字串長度(位元組長度)
	fmt.Println("Bytes:", len(username))
	// 取得 rune 集合長度
	fmt.Println("Runes:",  len([]rune(username)))

	// 用切片擷取字串的前 10 個元素, 理論上剛好到 Ü
	fmt.Println(string(username[:10]))
	fmt.Println(string([]rune(username)[:10]))
}