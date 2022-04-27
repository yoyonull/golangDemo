package main

import "fmt"

//方法一
func ReverseOne(s string) string {
	a := []rune(s)
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return string(a)
}

//方法二
func ReverseTwo(s string) string {
	r := []byte(s)
	for i := 0; i < len(s); i++ {
		r[i] = s[len(s)-1-i]
	}
	return string(r)
}

//方法三
func ReverseThree(s string) string {
	str := []int32(s)
	t := len(s)
	//t-1 代表对应是从0开始的
	for i := 0; i < t/2; i++ {
		str[i], str[t-1-i] = str[t-1-i], str[i]
	}
	return string(str)
}

func main() {
	a := "lele12121hue"
	c := ReverseThree(a)
	fmt.Println(c)
}
