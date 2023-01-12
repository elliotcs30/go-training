// 推斷型別發生問題:

package main

import "math/rand"

func main() {
	// var seed = 1234456789
	
	// 宣告得到的變數會出現問題
	// 因為 seed 需要的參數型別為 int64
	// rand.Seed(seed)

	// cannot use seed (variable of type int) as type int64 in argument to rand.Seed

	var seed int64 = 1234456789
	rand.Seed(seed)
}