// 安全走訪一個字串

package main

import "fmt"

func main() {
	logLevel := "デバッグ"
	for index, runeVal := range logLevel {
		// 印出索引及轉成字串的 rune 字元
		fmt.Println(index, string(runeVal))
	}
}