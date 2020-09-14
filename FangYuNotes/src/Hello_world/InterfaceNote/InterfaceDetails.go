package InterfaceNote

import "fmt"

//接口的注意事项和细节说明
//接口本身不能创建实例,但可以指向一个实现了该接口自定义类型的变量
//接口中所有的方法都没有方法体，即没有实现的方法
//自定义类型需要实现其中的所有方法，才叫自定义类型实现了接口
//一个自定义类型只有实现了某个接口，才能将自定义类型的实例（变量）赋值给接口类型
//只要是自定义类型，都可以实现接口，不仅仅是结构体
//一个自定义类型可以实现多个接口
//golang接口中不能有任何的变量
//接口直接可以继承，如果实现对应接口，需要将继承的接口也全部实现
//interface默认是一个指针，如果没有对interface初始化就使用，那么就会输出nil
//空接口{}没有实现任何方法，所以所有类型都实现了空接口
type AInterface interface {
	Say()
}
type BInterface interface {
	SayHello()
}
type Stu struct {
	Name string
}

func (Stu *Stu) Say() {
	fmt.Println(Stu.Name)
}
func (Stu *Stu) SayHello() {
	fmt.Println(Stu.Name, Stu.Name)
}

//type integer int
//func (i integer) Say()  {
//	fmt.Println("i say is,",i)
//}
func InterFaceDetails() {
	stu := Stu{Name: "Tom"}
	var a AInterface = &stu
	b := BInterface(&stu)
	a.Say()
	b.SayHello()
}
