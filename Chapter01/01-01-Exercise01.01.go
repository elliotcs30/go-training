// 隨機挑出介於 1-5 之間的數字
// 依照挑選出的數字作為顯示星號字元 * 的個數

// 宣告套件
package main

// 從套件匯入額外功能
import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	const num = 5

	// 亂數產生設定種子
	rand.Seed(time.Now().UnixNano())
	
	// 產生介於 0-4 之間的亂數, 再加上 1, 使結果落在 1-5
	index := rand.Intn(num) + 1

	// 利用字串重複函式 strings.Repeat(字元, 次數)
	// 用來產生出來的亂數作為顯示星號的個數
	stars := strings.Repeat("*", index)
	fmt.Println(stars)
}