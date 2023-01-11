// Example01:
// 輸出隨機字串

package main

// 從套件匯入額外的功能
import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

// 試先定義的訊息陣列 
var helloList = []string{
	"Hello, world",
	"你好, 世界",
	"こんにちは世界",
	"سلام دنیا",
	"ওহে বিশ্ব ",
	"Καλημέρα κόσμε",
}

func main() {
	rand.Seed(time.Now().UnixNano()) 	 // 利用當下時刻作為亂數種子
	index := rand.Intn(len(helloList)) // 以陣列長度當成產生亂數的範圍
	msg, err := hello(index)					 // 呼叫 hello 自訂函式並接收多個傳回值
	if err != nil {		// 若有錯誤，顯示錯誤並終止程式
		log.Fatal(err)
	}
	fmt.Println(msg)  // 沒有錯誤就將訊息顯示在主控台
}

// 自訂函式
func hello(index int) (string, error) {
	// 若收到的亂數超出 helloList 範圍
	// 產生一個錯誤並包含錯誤訊息，
	// 在訊息中把亂數從整數型別轉換成字串
	if index < 0 || index > len(helloList) - 1 {
		return "", errors.New("out of range: " + strconv.Itoa(index))
	}
	
	// 沒有錯誤就傳回訊息
	return helloList[index], nil
}
