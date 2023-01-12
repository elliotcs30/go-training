// 延伸習題: 定義與列印
// 建立一個醫療表格, 以便填寫患者的姓名、年齡、以及是否對花生過敏
// 1. 首先建立以下變數:
// 	* 名字為字串
// 	* 姓氏為字串
// 	* 年齡是整數 int
// 	* 過敏與否以布林值 bool 表示
// 2. 確保以上變數都以賦予初始值
// 3. 將資料值印到主控台畫面上
// Bob
// Smith
// 34
// false

package main

import (
	"fmt"
)

func main() {
	firstName := "Bob"
	familyName := "Smith"
	age := 34
	peanutAllergy := false

	fmt.Println(firstName)
	fmt.Println(familyName)
	fmt.Println(age)
	fmt.Println(peanutAllergy)
}