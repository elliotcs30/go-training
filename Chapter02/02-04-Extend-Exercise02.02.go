// 延伸習題
// 用 range 來走訪 map 裡的資料
// 假設有一下方的資料表, 找出出現次數最多的字, 並將其次數和字印出來

// ------------------
// |  Word  | Count |
// ------------------
// | Gonna  |   3   |
// | You    |   3   |
// | Give   |   2   |
// | Never  |   1   |
// | Up     |   4   |
// ------------------

// 將這些資料整理成一個 map 集合
// 用 for range 迴圈走訪它, 
// 紀錄出現次數最多的單字, 
// 以及該單字的出現次數, 然後印出以下結果:
// 出現最多次的字: Up
// 出現次數: 4

package main

import "fmt"

func main() {
	words := map[string]int {
		"Gonna": 3,
		"You": 	 3,
		"Give":  2,
		"Never": 1,
		"Up":    4,
	}

	topWord := ""
	topCount := 0
	for key, value := range words {
		if value > topCount {
			topCount = value
			topWord = key
		}
	}
	fmt.Println("出現最多次的字:", topWord)
	fmt.Println("出現次數:", topCount)
}
