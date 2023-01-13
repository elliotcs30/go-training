// 沒有運算式的 switch

package main

import (
	"fmt"
	"time"
)

func main() {
	// 只有起始賦值敘述
	switch dayBorn := time.Sunday; {
	case dayBorn == time.Sunday || dayBorn == time.Saturday:
		fmt.Println("生日為週末")
	default:
		fmt.Println("生日非週末")
	}
}
