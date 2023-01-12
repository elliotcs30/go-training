// 使用起始賦值敘述的 if 敘述

package main

import(
	"errors"
	"fmt"
)

// 會傳回 error 值的檢查函式
func validate(input int) error {
	if input < 0 {
		return errors.New("輸入值不得為負")
	} else if input > 100 {
		return errors.New("輸入值不得超過 100")
	} else if input%7 == 0 {
		return errors.New("輸入值不能為7的倍數")
	} else {
		return nil // 檢查都通過時傳回 nil
	}
}

func main() {
	input := 22
	
	// 接收 error 並檢查是否有錯誤
	if err := validate(input); err != nil {
		fmt.Println(err)
	} else if input%2 == 0 {
		fmt.Println(input, "是偶數")
	} else {
		fmt.Println(input, "是奇數")
	}
}

