package main

import "fmt"

func main() {
	//a:=make(map[string]int)
	//a["age"] =10

	var a = []int64{1, 2, 3}
	fmt.Printf("原始地址%p\n", &a)
	modifySlice(a)
	fmt.Printf("改动后地址%v\n", a)
}

func modifySlice(a []int64) {
	fmt.Printf("接收地址%p\n", &a)
	a[0] = 10
}
