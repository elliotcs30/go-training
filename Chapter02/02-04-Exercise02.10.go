// 利用迴圈走訪 map 元素

package main

import "fmt"

func main() {
	config := map[string]string {
		"debug": "1",
		"logLevel": "warn",
		"version": "1.2.1",
	}

	for key, value := range config {
		fmt.Println(key, "=", value)
	}
}