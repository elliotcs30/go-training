// 建立一個切片集合 (int, int8 型別)
// 在集合中放進一千萬個數字。
// 使用 runtime 套件取得整支程式所使用的堆積記憶體量(位元組)
// 轉換成百萬位元組單位後印出

package main

import (
	"fmt"
	"runtime"
)

func main() {
	//var list []int  // 469 MiB
	var list []int8   // 49 MiB
	for i := 0; i < 10000000; i++ {
		list = append(list, 100)
	}

	// 印出堆疊記憶體用量
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("TotalAlloc (heap) = %v MiB\n", m.TotalAlloc/1024/1024)
}