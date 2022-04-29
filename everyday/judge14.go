package main

import (
	"fmt"
)

func main() {

	abc := make(chan int, 10)
	timeout := make(chan bool)
	for i := 0; i < 10; i++ {
		abc <- i
	}
	go func() {
		for a := range abc {
			fmt.Println("a: ", a)
		}
		timeout <- true
	}()
	//借用select 关闭
	for {
		select {
		case <-timeout:
			close(abc)
		default:
			break
			//time.Sleep(time.Second * 3)
			//fmt.Println("close")
		}
	}
	//close(abc)
	//fmt.Println("close")
	//time.Sleep(time.Second * 10)
	fmt.Println("exit")

}
