// switch 敘述與多重 case 配對值
// 判斷生日當天是工作日還是週末, 並印出一段訊息

package main

import (
	"fmt"
	"time"
)

func main() {
	dayBorn := time.Sunday
	switch dayBorn {
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("生日為平日")
	case time.Saturday, time.Sunday:
		fmt.Println("生日為週末")
	default:
		fmt.Println("生日錯誤")
	}
}