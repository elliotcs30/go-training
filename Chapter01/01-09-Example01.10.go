// 變數遮蔽

package main

import "fmt"

var level = "pkg"

func main() {
	fmt.Println("Main start :", level)

	level := 42 // 在 main() 範圍定義 level
	if true {
		fmt.Println("Block start :", level)
		funcA()
	}
	fmt.Println("Main end :", level)
}

func funcA() {
	fmt.Println("funcA start :", level)
}