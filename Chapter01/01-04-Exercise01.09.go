// 算符(operators)
// 用算符處理數字

package main

import "fmt"

func main() {
	// 主餐 (2人, 每份 $13)
	var total float64 = 2 * 13
	fmt.Println("+ 主餐:", total)

	// 飲料 (每杯 $2.25, 4杯)
	total = total + (4 * 2.25)
	fmt.Println("+ 飲料:", total)

	// 折抵金額 ($5)
	total = total - 5
	fmt.Println("+ 折抵:", total)

	// 10% 服務費 (小費)
	tip := total * 0.1
	total = total + tip
	fmt.Println("+ 小費:", total)

	// 分攤額
	split := total / 2
	fmt.Println("分攤額:", split)

	// 來店次數
	visitCount := 24 // 過去來店 24 次
	visitCount = visitCount + 1
	
	// 用餘數算符檢查除以 5 餘數是否為 0 (是 5 的倍數), 傳回布林值
	remainder := visitCount % 5
	if remainder == 0 {
		fmt.Println("你已獲得來店滿 5 次折價券!")
	}
}
