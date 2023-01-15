// 溢位和越界繞回

package main

import "fmt"

func main() {
	var a int8 = 128
	fmt.Println(a)
}

// Terminal:
// cannot use 128 (untyped int constant) as int8 value in variable declaration (overflows)