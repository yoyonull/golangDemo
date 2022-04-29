package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	name string
}

func main() {
	//*Student Student
	m := map[string]*Student{"people": {"zhoujielun"}}
	m["people"].name = "wuyanzu"
	fmt.Printf("%v", m)
	fmt.Println("****")
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	//输出序列化后的结果
	fmt.Printf("a map 序列化后=%v\n", string(data))

}
