// 用 for 迴圈走訪切片元素

package main

import "fmt"

func main() {
	names := []string {"Elliot", "Joe", "Bonnie"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}