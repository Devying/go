package main

import "fmt"

func main() {
	//直接 slice:=make([]int,5,10)
	slice := []int{11, 22, 33, 44, 55, 66}
	//基于现有的切片或者数组创建，使用[i:j]这样的操作符即可，
	//她表示以i索引开始，到j索引结束,截取原数组或者切片，
	//创建而成的新切片，新切片的值包含原切片的i索引，但是不包含j索引
	slice1 := slice[:]
	slice2 := slice[0:]
	slice3 := slice[:5]
	slice4 := slice[3:5]
	newSlice := slice[1:3]
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)

	fmt.Println(newSlice)
	//计算长度
	fmt.Printf("slice4长度:%d,容量:%d\n", len(slice4), cap(slice4))
	fmt.Printf("newSlice长度:%d,容量:%d\n", len(newSlice), cap(newSlice))

	//追加
	newSlice = append(newSlice, 10, 20, 30, 40, 50, 60) //追加
	slice4 = append(slice4, newSlice...)
	fmt.Println(slice4)
	fmt.Println(newSlice)

	for _, v := range slice4 {
		fmt.Printf("值:%d\n", v)
	}
}
