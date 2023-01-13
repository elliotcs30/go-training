// break 和 continue 來控制迴圈
// 隨機產生介於 0-8 之間的數字,
// 迴圈要略過任何 3 的倍數, 若是偶數就結束迴圈

package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	for {
		rand.Seed(time.Now().UnixNano())
		r := rand.Intn(8)
		if r%3 == 0 {
			fmt.Println("略過")
			continue
		} else if r%2 == 0 {
			fmt.Println("跳出")
			break
		}
		fmt.Println(r)
	}
}