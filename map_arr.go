package main

import "fmt"

func main() {
	//实现一个php的关联数组多维的
	info := map[string]interface{}{"name": "huangby", "age": 28,"address":"Haidian,Beijing"}
	fmt.Println(info);

	info1 := make(map[string]interface{});
	info1["name"] =  "huangbaoying"
	info1["age"] =  28
	info1["job"] =  "developer"
	fmt.Println(info1);

	info2 := make(map[string]interface{});
	info2["name"] =  "humay"
	info2["age"] =  27
	info2["job"] =  "doctor"
	fmt.Println(info2);

	var user []map[string]interface{}
	user = append(user,info,info1,info2);
	fmt.Println(user);

}