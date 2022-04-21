package main

import "fmt"

func main() {

	//定义三个通道 1.写入 2.读出，判断 3.退出

	inchan := make(chan int, 1000)
	reschan := make(chan int, 2000)
	exitchan := make(chan bool, 4)

	go put(inchan)
	//todo定义少了
	for i := 0; i < 4; i++ {
		go getRes(inchan, reschan, exitchan)
	}

	go func() {
		for i := 0; i < 4; i++ {
			<-exitchan
		}
		close(reschan)
	}()

	//打印素
	for {
		rea, ok := <-reschan
		if !ok {
			break
		}
		fmt.Printf("%d\n", rea)

	}

}

func getRes(inchan chan int, reschan chan int, exitchan chan bool) {
	//素数处理
	var flag bool
	for {

		num, ok := <-inchan
		if !ok {
			break
		}
		flag = true //假设是素数
		for i := 2; i < num; i++ {
			if num%i == 0 { //合数
				flag = false
				break
			}
		}

		if flag {
			reschan <- num
		}

	}
	fmt.Println("有一个primeNum 协程因为取不到数据，退出")
	//这里我们还不能关闭 primeChan
	//向 exitChan 写入true
	exitchan <- true
}

func put(inchan chan int) {
	for i := 1; i <= 8000; i++ {
		inchan <- i
	}
	//关闭
	close(inchan)
}
