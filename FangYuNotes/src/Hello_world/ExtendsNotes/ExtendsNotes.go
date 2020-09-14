package ExtendsNotes

import "fmt"

//golang 中使用的是匿名结构体进行继承
//
type Stu struct {
	Name  string
	Age   int
	Score int
}

func (stu *Stu) ShowInfo() {
	fmt.Println(stu)
}

func (stu *Stu) SetScore(score int) {
	stu.Score = score
}

func (stu *Stu) testing() {
	fmt.Println("Testing .......")
}

//继承
type Graduate struct {
	Stu //嵌入了Stu的匿名结构体
	GPA float64
}

func ExtendsNotes() {
	graduate := &Graduate{}
	graduate.Name = "Tom"
	graduate.Age = 20
	graduate.GPA = 3.7
	graduate.testing()
	graduate.SetScore(90)
	graduate.ShowInfo()
}
