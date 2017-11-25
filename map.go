package main

import "fmt"

func main() {
	dict := map[string]int{"张三": 43, "李四": 50}
	fmt.Println(dict)

	dict1 := map[string]string{}
	dict1["name"] = "huangby"
	fmt.Println(dict1)
}
