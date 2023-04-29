package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	jsonBuf := `
    {
    "company":"Google",
    "subjects":[
            "Go",
            "C++",
            "C",
            "PHP"
    ],
    "isok":true,
    "price":3.1415,
	"colors":""
    }
    `
	ujson := make(map[string]interface{}, 4)
	err := json.Unmarshal([]byte(jsonBuf), &ujson)
	if err != nil {
		fmt.Printf("err is %s\n", err)
	}
	//通过反射获取对应的信息
	for k, v := range ujson {
		switch data := v.(type) {
		case string:
			fmt.Printf("map[%s]的值类型为string, value = %s\n", k, v)
		case bool:
			fmt.Printf("map[%s]bool, value = %v\n", k, data)
		case float64:
			fmt.Printf("map[%s]float64, value = %v\n", k, data)
		case []interface{}: //利用万能指针 不能直接判断切片
			fmt.Printf("map[%s]interface{}, value = %v\n", k, data)
		}

	}

}
