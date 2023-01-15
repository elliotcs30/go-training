// 浮點數的精確度除3再乘3

package main

import "fmt"

func main() {
	var a int = 100
	var b float32 = 100
	var c float64 = 100

	fmt.Println((a / 3) * 3) // 99
	fmt.Println((b / 3) * 3) // 100
	fmt.Println((c / 3) * 3) // 100
}