// 延伸習題: 指標值替換
// 補足漏缺的程式碼, 用註解提供說明, 
// 而且必須讓變數 a 和 b 的值交換
// swap() 函式只能接收指標, 而不會傳回任何值

package main

import "fmt"

func main() {
	a, b := 5, 10

	// 在此呼叫交換函式
	swap(&a, &b)
	fmt.Println(a == 10, b == 5)
}

func swap(a *int, b *int) {
	
	// 在此交換變數值
	*a, *b = *b, *a
}

// true true


