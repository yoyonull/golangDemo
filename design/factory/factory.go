package main

import (
	"fmt"
	"golangDemo/design/factory/model"
)

func main() {

	var stu = model.NewStudent("chen", 98.5, 30)
	fmt.Println(stu) //带有指针符号&,&{chen 98.5 30}

	fmt.Println("name=", stu.Name, "age=", stu.GetAge())
	//*stu.Name 是一样的，编译器自动给处理了（获取stu指针的属性）
}
