package main

import (
	"fmt"
)

type Person struct {
	name   string
	gender string
	age    int
}

func (p *Person) SayHi() {
	fmt.Printf("%v:%v:%v\n", p.name, p.gender, p.age)
}

type Student struct {
	school string
	Person //匿名字段
}
type Employee struct {
	company string
	Person  //匿名字段
}

func main() {
	//c:=&Person{"lxx", "male", 18}
	//fmt.Println(c)
	//fmt.Println(&c)
	//fmt.Println(*c)
	t := make(chan int, 2)

	//for  {
	//fmt.Println(111)
	select {
	case t <- 000:
		fmt.Println(t)
		fmt.Println(110)
	default:
		fmt.Println(220)

	}
	//}

	s1 := Student{"佳丽蹲", Person{"lxx", "male", 18}}
	e1 := Employee{"智力有限公司", Person{"egon", "male", 18}}
	s1.SayHi()
	e1.SayHi()
}
