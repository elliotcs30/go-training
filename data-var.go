
// 資料型態:
// 整數 int(-32, 51120)
// 浮點數 float(3.1415926)
// 字串 string("Hello Golang")
// 布林值 bool(true、false)
// 字元 rune('a'), 背後實際存放整數 a = 97

// 變數: 可以存放資料的空間
// 宣告變數 > 指定資料 > 使用變數
// 宣告變數: var 變數名稱 資料型態名稱
// 指定資料: 變數名稱 = 對應型態的資料
// 使用變數: 使用變數名稱代替資料來做運算

package main // 可執行程式必須使用 main 組件
import "fmt" // 載入內建的 fmt 封包，用來做基本輸入
func main() { // 建立 main 函式，程式的進入點
	
	// 執行程式 start
	/*
	fmt.Println(3) // 整數 int
	fmt.Println(3.1415926) // 浮點數 float64
	fmt.Println("測試中文") // 字串 string
	fmt.Println(true) // 布林值 bool
	fmt.Println('a')	// 字符 rune
	*/

	// 變數的使用
	var x int 			// 宣告變數 x
	x = 5 					// 指定資料: 把資料 5 放進變數裡
	fmt.Println(x) 	// 使用變數: 用變數的名稱代替資料，做操作
	x = 10					// 指定新的資料: 把資料 10 放進變數，覆蓋原來的資料
	fmt.Println(x)

	var f float64 = 3.1415 // 宣告變數順便放進資料: var 變數名稱 資料型態 = 適當的資料
	fmt.Println(f)

	var s string = "Hello Golang"
	fmt.Println(s)

	var test bool = true
	fmt.Println(test)

	var c rune = 'e'
	fmt.Println(c)
}

// terminal go run data-var