// Rune

package main

import "fmt"

func main() {
	username := "Sir_King_Über"

	// 走訪字串中的每一個字元
	for i := 0; i < len(username); i++ {
		// 印出一個字元加一個空格
		fmt.Print(string(username[i]), " ") // S i r _ K i n g _ Ã  b e r
	}
}