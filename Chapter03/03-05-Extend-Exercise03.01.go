// 延伸習題: 銷售計算機
// 請寫一個專供購物車應用程式使用的工具, 計算稅金並加總售價

// 1. 建立一個計算單品稅金的程式
// 2. 以售價和稅率作為計算機程式的輸入值
// 3. 把稅金加總, 在印出下表的稅金總額

// 購物清單及稅率:
// ---------------------
// | 項目 | 成本  | 稅率 |
// ---------------------
// | 蛋糕 | 0.99 | 7.5% |
// | 牛奶 | 2.75 | 1.5% |
// | 奶油 | 0.87 |   2% |
// ---------------------

// 輸出結果:
// Sales Tax Total: 0.1329

package main

import "fmt"

func salesTax(cost float64, texRate float64) float64 {
	return cost * texRate
}

func main() {
	taxTotal := .0
	
	taxTotal += salesTax(.99, .075)
	taxTotal += salesTax(2.75, .015)
	taxTotal += salesTax(.87, .02)
	fmt.Println("Sales Tax Total: ", taxTotal)
}
