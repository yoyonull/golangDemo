package main

import "fmt"

func main() {

	//todo go 读写竞争快速搞
	i := 0

	go func() {
		i++
	}()
	fmt.Println(i)
}
