// 從子範圍存取上層變數

package main

import "fmt"

var level = "pkg"

func main() {
	// main() 層級
	fmt.Println("Main start :", level)

	if true {
		// main() 底下的 if 層級
		fmt.Println("Block start :", level)
		funcA()
	}
}

func funcA() {
	// funcA() 函式層級
	fmt.Println("funcA start :", level)
}