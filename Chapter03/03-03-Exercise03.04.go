// 使用大數值

package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	intA := math.MaxInt64
	intA = intA + 1

	bigA := big.NewInt(math.MaxInt64) // big.Int 整數
	bigA.Add(bigA, big.NewInt(1))

	fmt.Println("MaxInt64: ", math.MaxInt64)
	fmt.Println("Int:     ", intA)
	fmt.Println("Big Int:  ", bigA.String()) // 將 big.Int 轉成字串印出
}

// MaxInt64:  9223372036854775807 -> int64 最大值
// Int:      -9223372036854775808 -> int 發生越界繞回
// Big Int:   9223372036854775808 -> big int 正確加 1