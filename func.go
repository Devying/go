package main

import (
	"fmt"
)

func main() {
	sum := add(1, 2)
	fmt.Println(sum)
}

// 参数类型  返回值类型

// 如果add 首字母大写是可以被其他地方引入包后使用的 比如Add
func add(a, b int) int {
	return a + b
}
