// 宣告常數
// 問題: 資料庫伺服器運算提升, 需要建立一個自訂的記憶體快取
// 使用 map 集合型別, 來擔任快取表, 但這個快取表中可容納的總項目數量是有上限的
// 快取中需要儲存兩種類型的資料: 書本和 CD 唱片
// 兩者都有識別用的 ID 字串和對應的名稱, 因此需要想辦法在共用的快取表中區分這兩種資料, 並有辦法讀取和寫入資料

package main

import "fmt"

// 單筆資料上限
const GlobalLimit = 100
// 快取最大容量 (單筆上限 x 2)
const MaxCacheSize int = 2 * GlobalLimit

const (
	// 書本 id 的前綴字
	CacheKeyBook = "book_"
	
	// CD id 的前綴字
	CacheKeyCD = "cd_"
)

// 快取集合
var cache map[string]string 

// 從快取取出某個鍵的值
func cacheGet(key string) string {
	return cache[key]
}

// 將某個鍵和值寫入快取
func cacheSet(key, val string) {
	// 如果快取大小已達極限就跳出函式
	if len(cache) + 1 >= MaxCacheSize {
		return
	}

	// 寫入資料
	cache[key] = val
}

// 加入書本資料
func SetBook(isbn string, name string) {
	// 加上書本前綴字後呼叫 cacheSet()
	cacheSet(CacheKeyBook + isbn, name)
}

// 讀取書本資料
func GetBook(isbn string) string {
	// 加上書本前綴字後呼叫 cacheGet()
	return cacheGet(CacheKeyBook + isbn)
}

// 寫入 CD 資料
func SetCD(sku string, title string) {
	// 加上 CD 前綴字後呼叫 cacheSet()
	cacheSet(CacheKeyCD + sku, title)
}

// 讀取 CD 資料
func GetCD(sku string) string {
	// 加上 CD 前綴字後呼叫 cacheGet()
	return cacheGet(CacheKeyCD + sku)
}

func main() {
	// 初始化快取
	cache = make(map[string]string)

	// 在快取寫入資料
	SetBook("1234-5678", "Get Ready To Go")
	SetCD("1234-5678", "Get Ready To Go Audio Book")

	// 讀取和印出快取資料
	fmt.Println("Book :", GetBook("1234-5678"))
	fmt.Println("CD   :", GetCD("1234-5678"))
}