package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//对接口sort 很好的实现(参考例子)
//对于该接口，是对照文档sort interface 来实现的
//很好的一个官方案例
//实现三个方法就能实现该接口(Len,Less，Swap)

type Hero struct {
	Name string
	Age  int
}

type HeroSlice []Hero

//排序实现源码开放的接口
func (hs HeroSlice) Len() int {
	return len(hs)
}

//按照对应的顺序来排列 从小到大
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
}
func (hs HeroSlice) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}

func main() {
	//单个切片
	var intSlice = []int{0, 2312, 2, 5, 3, 6, 23, 8}
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	//i :=11
	//j :=22
	//i,j = j,i
	//fmt.Println("i=",i,"j=",j)

	//切片排序
	var heroes HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			//Name: fmt.Println("英雄"),
			//Name: fmt.Sprintf("英雄|#{rand.Intn(100)}"),
			Name: fmt.Sprintf("英雄|%d", rand.Intn(100)),

			Age: rand.Intn(100),
		}
		heroes = append(heroes, hero)
	}

	//排序前
	for _, v := range heroes {
		fmt.Println(v)
	}

	//排序来看
	sort.Sort(heroes)
	fmt.Println("----排序后----")
	for _, v := range heroes {
		fmt.Println(v)
	}

}
