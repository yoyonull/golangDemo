package main

import (
	"encoding/json"
	"fmt"
)

// Actress3 女演员
type Actress3 struct {
	Name       string
	Birthday   string
	BirthPlace string
	Opus       []string
}

func main() {

	// 普通JSON
	// 因为json.UnMarshal() 函数接收的参数是字节切片，   // 所以需要把JSON字符串转换成字节切片。
	jsonData := []byte(`{
	"name":"迪丽热巴",
		"birthday":"1992-06-03",
		"birthPlace":"新疆乌鲁木齐市",
		"opus":[
			"《阿娜尔罕》",
			"《逆光之恋》",
			"《克拉恋人》"
			]
    }`)

	var Actress3 Actress3
	err := json.Unmarshal(jsonData, &Actress3)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("姓名：%s\n", Actress3.Name)
	fmt.Printf("生日：%s\n", Actress3.Birthday)
	fmt.Printf("出生地：%s\n", Actress3.BirthPlace)
	fmt.Println("作品：")
	for _, val := range Actress3.Opus {
		fmt.Println("\t", val)
	}
}
