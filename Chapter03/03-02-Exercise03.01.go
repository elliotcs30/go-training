// 用程式判斷密碼複雜度
// 1. 必須有小寫
// 2. 必須有大寫字母
// 3. 必須有數字
// 4. 必須有符號
// 5. 長度必須至少有 8 個字元

package main

import (
	"fmt"
	"unicode" // 使用 unicode 函式庫
)

func passwordChecker(pw string) bool {
	// 把密碼轉成 rune 型別, 以便安全接收 UTT-8 字串
	pwR := []rune(pw)
	// 若密碼長度不足 8 字元, 等於檢查失敗
	if len(pwR) < 8 {
		return false
	}

	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSymbol := false

	// 用 for range 走訪字串的每個字元, 忽略其索引
	for _, v := range pwR {
		// 是否有大寫字元
		if unicode.IsUpper(v) {
			hasUpper = true
		}
		// 是否有小寫字元
		if unicode.IsLower(v) {
			hasLower = true
		}
		// 是否有數字
		hasNumber = unicode.IsNumber(v)
		// 是否有(標點符號)
		hasSymbol = unicode.IsPunct(v) || unicode.IsSymbol(v)
	}
	return hasUpper && hasLower && hasNumber && hasSymbol
}

func main() {
	if passwordChecker("") {
		fmt.Println("密碼格式良好")
	} else {
		fmt.Println("密碼格式不正確")
	}

	if passwordChecker("This!I5A") {
		fmt.Println("密碼格式良好")
	} else {
		fmt.Println("密碼格式不正確")
	}
}