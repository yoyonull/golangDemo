package main

import "fmt"

type Monkey struct {
	Name string
}

//定义接口扩展
type Bird interface {
	Fly()
}

//给monkey 绑定一个方法
func (this Monkey) climbing() {
	fmt.Println(this.Name, "会爬树")
}

type Smallmonkey struct {
	Monkey
}

func (this Smallmonkey) Fly() {
	fmt.Println(this.Name, "会飞行")
}

func main() {
	mm := Smallmonkey{
		Monkey{Name: "悟空"},
	}

	mm.climbing()

	mm.Fly()

}
