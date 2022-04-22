package model

type student struct {
	Name  string
	Score float64
	age   int
}

//因为student结构体首字母是小写，因此能在model中使用，若在其他地方使用
//需要创建方法使用工厂模式来解决

func NewStudent(n string, s float64, a int) *student {
	return &student{
		Name:  n,
		Score: s,
		age:   a,
	}
}

//如果结构体里面字段小写后，需要绑定改变
func (s *student) GetAge() int {
	return s.age
}
