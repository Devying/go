package main

import "fmt"

func main() {
	array := [5]int{1: 1, 3: 4} //只初始化索引位1和3的值
	for i, v := range array {
		fmt.Printf("索引:%d,值:%d\n", i, v)
	}
	//修改数组元素
	array[3] = 100

	for i, v := range array {
		fmt.Printf("索引:%d,值:%d\n", i, v)
	}

	//指针数组
	parray := [5]*int{1: new(int), 4: new(int)}

	//修改
	*parray[1] = 88
	*parray[4] = 99
	fmt.Printf("索引:1,值:%d\n", *parray[1])
	fmt.Printf("索引:4,值:%d\n", *parray[4])
}
