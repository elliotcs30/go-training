// 字串與字串常值

package main

import "fmt"

func main() {
	// 原始字串, 直接換行
	comment1 := `This is the BEST
	thing ever!`
	comment2 := `This is the BEST\nthing ever!` // 原始字串, 換行符號不會轉譯
	comment3 := "This is the BEST\nthing ever!" // 轉譯字串, 換行符號會轉譯

	fmt.Print(comment1, "\n\n")
	fmt.Print(comment2, "\n\n")
	fmt.Print(comment3, "\n")
}