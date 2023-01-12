// 子範圍的變數在外部無法取得
package main

import "fmt"

func main() {
	{
		level := "Nest 1"
		fmt.Println("Block end :", level)
	}

	// 產生錯誤: undefined: level, 因爲無法取得上面 {} 中的變數
	fmt.Println("Main end :", level)
}

// undefined: level