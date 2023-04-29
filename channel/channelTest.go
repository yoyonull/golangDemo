package main

import (
	"fmt"
)

func main() {

	//for  {
	//	//每隔1s 打印
	//	time.Sleep(2 * time.Second)
	//	fmt.Println(1111)
	//}

	chInt := make(chan int, 10)
	for i := 0; i < 10; i++ {
		chInt <- i
	}
	chString := make(chan string, 20)
	for i := 0; i < 20; i++ {
		chString <- "hello" + fmt.Sprintf("%d", i)
	}

	//for {
	//	time.Sleep(1 * time.Second)
	//	fmt.Println(<- chInt)
	//	fmt.Println(<- chString)
	//}
	////传统的方法在遍历管道时，如果不关闭会阻塞而导致 deadlock
	////使用的对应的
	for {
		select {
		case i := <-chInt:
			fmt.Printf("chanInt%d\n", i)
		case i := <-chString:
			fmt.Printf("chString%s\n", i)
		default:
			fmt.Printf("都取不到了，不玩了, 程序员可以加入逻辑\n")
			return
		}

	}

}
