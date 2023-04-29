package main

import "fmt"

func main() {
	a := make(map[string]int)
	a["age"] = 10
	fmt.Printf("原始地址%p\n", &a)
	modify(a)
	fmt.Printf("改动后地址%v\n", a)
}

func modify(a map[string]int) {
	fmt.Printf("接收地址%p\n", &a)
	a["age"] = 1
}
