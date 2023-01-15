// Rune

package main

import "fmt"

func main() {
	username := "Sir_King_Über"

	// 走訪字串中的每一個字元
	for i := 0; i < len(username); i++ {
		// 印出一個字元加一個空格
		fmt.Print(username[i], " ") // 83 105 114 95 75 105 110 103 95 195 156 98 101 114
	}
}