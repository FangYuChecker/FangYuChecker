package ExtendsNotes

import "fmt"

type Monkey struct {
	Name string
}

type Bird interface {
	Flying()
}

type Fish interface {
	Swimming()
}

type Lesson interface {
	Bird
	Fish
}

func (this *Monkey) Climbing() {
	fmt.Println(this.Name, "生来会爬树.....")
}

func (this *LittleMonkey) Flying() {
	fmt.Println(this.Name, "学会了飞翔.....")
}

func (this *LittleMonkey) Swimming() {
	fmt.Println(this.Name, "学会了游泳.....")
}

type LittleMonkey struct {
	Monkey //extend
}

func IFANDET() {
	monkey := LittleMonkey{
		Monkey{
			Name: "Tom",
		},
	}
	monkey.Climbing()
	monkey.Flying()
	monkey.Swimming()
}

//实现接口可以看作对继承的一种补充
//继承的价值主要体现在复用性和可维护性
//接口的价值在于设计好规范，让其他自定义类型去实现这些方法
//接口比继承更加灵活
//接口在一定程度上实现代码解耦
