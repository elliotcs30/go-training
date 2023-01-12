// 運用簡寫算符

package main 

import "fmt"

func main() {
	// 宣告一個變數並賦予初始值
	count := 5

	// 變數加上 5、再重新賦值回給同一變數
	count += 5
	fmt.Println(count)

	// 變數遞增 1
	count++
	fmt.Println(count)

	// 變數遞減 1
	count--
	fmt.Println(count)

	// 變數減去 5
	count -= 5
	fmt.Println(count)

	name := "John"
	name += " Smith" // 對原本的字串串接字串
	fmt.Println("Hello, ", name)
}