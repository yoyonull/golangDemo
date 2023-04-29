package main

import (
	"errors"
	"fmt"
)

//todo defer recover
func test() {
	defer func() {
		err := recover() //recover是内置函数 ，可以捕获异常
		if err != nil {
			errors.New("错误说明")
			fmt.Println("err=", err)
		}
	}()
	a := 10
	b := 0
	res := a / b
	fmt.Println(res)

}
func test2() {
	defer func() {
		err := recover() //recover是内置函数 ，可以捕获异常
		if err != nil {
			fmt.Println("err is ", err)
		}
	}()
	a := 10
	b := 0
	res := a / b
	errors.New("12121212")
	fmt.Println(res)

}

func readConf(name string) (err error) {
	if name == "config.ini" {
		//读取...
		return nil
	} else {
		//返回一个自定义错误
		return errors.New("读取文件错误..")
	}
}

func main() {
	test()
	err := readConf("config2.ini")
	if err != nil {
		//如果读取文件发送错误，就输出这个错误，并终止程序
		panic(err)
	}
	fmt.Println("继续执行....")

	//arr1 := [...]int{1, 3, 5, 7, 9}
	//s1 := arr1[1:3]
	//fmt.Println(s1)

	var arr = [4][2]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
	c := 1
	for i, v := range arr {
		fmt.Println(i, v, c)
	}

}
