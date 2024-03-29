package main

import (
	"encoding/json"
	"fmt"
)

/*{"addr":"beijing","age":123,"name":"typ"}
{"name":"qqqq","age":12,"hobbit":"ttttt"}
[{"addr":"beijing111","age":1231,"name":"typ111"},{"addr":"beijing222","age":1222,"name":"typ222"}]*/

type student struct {
	//tagjson序列化后是小写
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Hobbit string `json:"hobbit"`
}

func unjsonStruct() {
	var stu student
	fmt.Println("unjson before： ", stu)
	var str = "{\"name\":\"qqqq\",\"age\":12,\"hobbit\":\"ttttt\"}"
	err := json.Unmarshal([]byte(str), &stu)
	if err != nil {
		fmt.Println("unjson 失败")
	}
	fmt.Println("unjson after： ", stu)
}

func unjsonMap() {
	testMap := make(map[string]interface{})
	fmt.Println("unjson before： ", testMap)
	var str = "{\"addr\":\"beijing\",\"age\":123,\"name\":\"type\"}"
	err := json.Unmarshal([]byte(str), &testMap)
	if err != nil {
		fmt.Println("unjson 失败")
	}
	fmt.Println("unjson after： ", testMap)
	for index, v := range testMap {
		fmt.Printf("\t日期：%s -- %s \n", index, v)
		//fmt.Printf("\t\t分类：%s\n", value.Type)
		//fmt.Printf("\t\t标题：%s\n", value.Title)
	}
}

func unjsonSlice() {
	var sli = make([]map[string]interface{}, 2)
	fmt.Println("unjson before： ", sli)
	var str = "[{\"addr\":\"beijing111\",\"age\":1231,\"name\":\"typ111\"},{\"addr\":\"beijing222\",\"age\":1222," +
		"\"name\":\"typ222\"}]"
	err := json.Unmarshal([]byte(str), &sli)
	if err != nil {
		fmt.Println("unjson 失败")
	}
	fmt.Println("unjson after： ", sli)
	for index, v := range sli {
		fmt.Printf("\t信息：%d -- %s \n", index, v)
		//fmt.Printf("\t\t分类：%s\n", value.Type)
		//fmt.Printf("\t\t标题：%s\n", value.Title)
	}
}

func main() {
	//unjsonStruct()
	//unjsonMap()
	unjsonSlice()
}
