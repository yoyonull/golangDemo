package main

import (
	"encoding/json"
	"fmt"
)

type Opus2 struct {
	Date  string
	Title string
}
type Actress2 struct {
	Name       string
	Birthday   string
	BirthPlace string
	Opus2      []Opus2
}

func main() {
	// JSON嵌套数组JSON
	jsonData := []byte(`{
      "name":"迪丽热巴",
      "birthday":"1992-06-03",
      "birthPlace":"新疆乌鲁木齐市",
      "Opus2":[
         {
            "date":"2013",
            "title":"《阿娜尔罕》"
         },
         {
            "date":"2014",
            "title":"《逆光之恋》"
         },
         {
            "date":"2015",
            "title":"《克拉恋人》"
         }
      ]
   }`)

	var Actress2 Actress2
	err := json.Unmarshal(jsonData, &Actress2)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("姓名：%s\n", Actress2.Name)
	fmt.Printf("生日：%s\n", Actress2.Birthday)
	fmt.Printf("出生地：%s\n", Actress2.BirthPlace)
	fmt.Println("作品：")
	for _, val := range Actress2.Opus2 {
		fmt.Printf("\t%s - %s\n", val.Date, val.Title)
	}
	for _, val := range Actress2.Opus2 {
		fmt.Printf("%s", val)
	}
}
