// Exercise01
// 隨機挑出一個介於 1-5 之間的數字
// 然後以這個數字作為在主控台顯示星星字元(*)的個數。

package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(5) + 1
	stars := strings.Repeat("*", index)
	fmt.Println(stars)
}
