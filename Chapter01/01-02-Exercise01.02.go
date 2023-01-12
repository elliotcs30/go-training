// 練習: 用 var 宣告一個變數

// 定義套件名稱
package main 

// 匯入套件
import (
	"fmt"
)

// 宣告套件範圍變數
var foo string = "bar"

func main() {
	// 宣告函式範圍變數
	var baz string = "qux"
	// 印出變數
	fmt.Println(foo, baz)
}
