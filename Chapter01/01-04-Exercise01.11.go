// 值的比較:
// 依照顧客的來店次數決定會員等級
// * 銀牌會員: 來店次數介於 11 - 20 次
// * 金牌會員: 來店次數介於 21 - 30 次
// * 白金 VIP: 來店次數超過 30 次

package main

import "fmt"

func main() {
	// 顧客目前來店的次數
	visits := 15

	// 判斷顧客屬於哪種等級的會員
	fmt.Println("新顧客 :", visits == 1)
	fmt.Println("熟客 :", visits != 1)
	fmt.Println("銀牌會員 :", visits > 10 && visits <= 20)
	fmt.Println("金牌會員 :", visits > 21 && visits <= 30)
	fmt.Println("白金 VIP :", visits > 30)
}