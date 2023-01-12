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

// 事先定義的訊息陣列
var helloList = []string{
	"Hello, world",
	"Καλημέρα κόσμε",
	"こんにちは世界",
	" ایند مالس",
	"Привет, мир",
}

// 主函式
func main() {
	// 抓取當前時間作為亂數種子
	rand.Seed(time.Now().UnixNano())

	// 以陣列長度當成產生亂數的範圍
	index := rand.Intn(len(helloList))

	// 呼叫 hello 函式並接收多個傳回值
	msg, err := hello(index)

	// 若有錯誤就顯示錯誤並終止程式
	if err != nil {
		log.Fatal(err)
	}
	// 沒有錯誤就將訊息顯示在主控台
	fmt.Println(msg)
}

// hello 函式
func hello(index int) (string, error) {
	if index < 0 || index > len(helloList)-1 {
		// 若收到的亂數超出 helloList 範圍, 產生一個錯誤並包含錯誤訊息,
		// 在訊息中把亂數從整數型別轉換成字串
		return "", errors.New("out of range: " + strconv.Itoa(index))
	}
	return helloList[index], nil //沒有錯誤就傳回訊息
}