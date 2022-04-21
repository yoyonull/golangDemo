package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	tt := []byte("this is test")

	fmt.Println(tt)                 //[116 104 105 115 32 105 115 32 116 101 115 116]
	fmt.Sprintf("this is %s", tt)   // 空空如也 wuio输出
	fmt.Printf("this iss %s\n", tt) //this iss this is test

	//还有中转换
	strtt := fmt.Sprintf("%s", tt)
	fmt.Println(strtt)

	//定义 数组map
	var slice []map[string]interface{}
	a := make(map[string]interface{})
	a["name"] = "chen"
	a["sex"] = "male"
	a["age"] = 1
	slice = append(slice, a)
	b := make(map[string]interface{})
	b["name"] = "高"
	b["sex"] = "female"
	b["age"] = 11
	slice = append(slice, b)

	data, _ := json.Marshal(slice)
	fmt.Println(data)
	fmt.Printf("slice序列化的结果%v\n", string(data))
	//dd :="[{\"age\":1,\"name\":\"chen\",\"sex\":\8"male\"},{\"age\":11,\"name\":\"高\",\"sex\":\"female\"}]\n"
	json.Unmarshal([]byte(data), &slice)
	//fmt.Println(data)
	fmt.Printf("反序列化的俄结果%v\n", string(data))

}
