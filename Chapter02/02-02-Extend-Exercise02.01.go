// 延伸習題: 實作 FizzBuzz
// 1. 寫一段可以從 1 顯示到 100 的程式
// 2. 如果顯示的數字是 3 的倍數, 改為顯示文字 Fizz
// 3. 如果顯示的數字是 5 的倍數, 改為顯示 Buzz
// 4. 如果顯示的數字是 3、5 的公倍數, 改為顯示 FizzBuzz

// 提示:
// 可以用 strconv.Itoa() 函式將數字轉換成字串
// 寫一個可以迭代 100 次的迴圈, 用一個變數來追蹤檢查的數字。第一個數字是 1, 最後一個是 100。使用上述的計數變數和餘數除法算符 %, 檢查是否為 3 或 5 的倍數。

package main

import (
	"fmt"
	// "strconv"
)

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 & i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else {
			fmt.Println(i)
		}
	}
}

// 使用 strconv
// func main() {
// 	for i := 1; i <= 100; i++ {
// 		out := ""
// 		if i%3 == 0 {
// 			out += "Fizz"
// 		}
// 		if i%5 == 0 {
// 			out += "Buzz"
// 		}
// 		if out == "" {
// 			out = strconv.Itoa(i)
// 		}
// 		fmt.Println(out)
// 	}
// }
