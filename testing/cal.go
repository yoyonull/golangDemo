package cal

//一个被测试函数
func addUpper(n int) int {
	res := 0
	for i := 1; i <= n-1; i++ {
		res += i
	}
	return res
}

//求两个数的查
func getSub(n1 int, n2 int) int {
	return n1 - n2
}

//测试方法
//go test -v 一个文件(同时 -v 包含测试的细节)
//如 go test -v test.run TestReStore 其中一个方法
