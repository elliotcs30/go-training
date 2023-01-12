// 在函式使用指標參數

package main

import "fmt"

// 傳入值
func add5Value(count int) {
	count += 5
	fmt.Println("add5Value: ", count)
}

// 傳入指標
func add5Point(count *int) {
	*count += 5
	fmt.Println("add5Point: ", *count)
}

func main() {
	var count int
	add5Value(count)
	fmt.Println("add5Value post: ", count)
	add5Point(&count)
	fmt.Println("add5Point post: ", count)
}

// add5Value     :  5
// add5Value post:  0
// add5Point     :  5
// add5Point post:  5