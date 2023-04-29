package main

import "fmt"

type Cat1 struct {
	name string
	age  int
}

func main() {

	testAllType()

	//testMap()
	//testSlice()

}

func testSlice() {
	//var mapchan chan Cat1
	mapchan := make(chan Cat1, 10)
	cat1 := Cat1{name: "一万", age: 1}
	cat2 := Cat1{name: "6万", age: 6}
	mapchan <- cat1
	mapchan <- cat2
	c := <-mapchan
	fmt.Println(c.name)
	//for k, v := range c {
	//	fmt.Sprintf("k is %s,v is %d\n",k,v)
	//}
	fmt.Println(c)

}

func testMap() {
	//var mapchan chan map[string]string
	mapchan := make(chan map[string]string, 10)
	//todo 注意前后的对应书写格式
	m1 := make(map[string]string, 10)
	m1["provice"] = "广东省"
	m1["city"] = "广州市"
	m2 := make(map[string]string, 10)
	m2["provice"] = "广东省"
	m2["city"] = "广州市"
	//
	mapchan <- m1
	mapchan <- m2
	cc := <-mapchan
	fmt.Println(cc)

}

func testAllType() {
	type1 := make(chan interface{}, 5)
	defer close(type1)
	type1 <- 10
	type1 <- "this"
	cat := Cat1{"小花猫", 15}
	type1 <- cat

	<-type1
	<-type1
	tt := <-type1
	fmt.Println(tt)
	fmt.Println(1111)
	//下面的写法是错误的!编译不通过
	//fmt.Printf("tt.Name=%v", tt.Name)
	//使用类型断言
	a := tt.(Cat1)
	fmt.Printf("tt.Name=%v/n", a.name)

}
