package main

import (
	"encoding/json"
	"fmt"
)

// Opus4 作品

type ColorGroup struct {
	ID     int      `json:"id"`
	Name   string   `json:"name"`   //todo Name 小写是不可以导出的
	Colors []string `json:"colors"` //tag 转化为小写
}

//type Group struct {
//	ID     int `json:"id"`
//	Name   string `json:"name"` //todo Name 小写是不可以导出的
//	Colors []string `json:"colors"` //tag 转化为小写
//}

func main() {
	cc := ColorGroup{
		ID:     1,
		Name:   "chen",
		Colors: []string{"black", "red", "white"},
	}

	c4, _ := json.Marshal(cc)

	fmt.Println(string(c4))
	var tt ColorGroup
	_ = json.Unmarshal(c4, &tt)
	//fmt.Println(string(c4))
	fmt.Printf("%#v\n", tt)
	fmt.Println(tt.Name)
	fmt.Println(tt.ID)
	fmt.Println(tt.Colors)
	for k, v := range tt.Colors {
		fmt.Printf("k is %d,v is %s\n", k, v)
	}

}
