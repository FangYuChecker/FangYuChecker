package Abstract

import "fmt"

//抽象出的字段和对字段的操作封装到一起，数据被保护在内部，程序只能通过被授权的操作才能对字段进行操作

//封装好处
//隐藏实现细节
//可以对数据进行验证，保证安全合理

//如何实现封装
//对结构体中的属性进行封装
//通过方法，包，实现封装

//如何实现封装
//将结构体，字段的首字母小写（即不能导出）
//将结构体提供一个工厂函数的形式，首字母大写，类似于一个构造函数
//提供一个首字母大写的Set方法，用于对属性判断并赋值
//提供一个首字母大写的Get方法，用于获取属性的value

type person struct {
	Name   string
	age    int //other package can not use
	salary float64
}

func Newperson(name string) *person {
	return &person{
		Name: name,
	}
}

//为了访问age个salart，我们编写set和get方法
func (p *person) SetAge(n int) {
	if n <= 0 || n > 150 {
		fmt.Println("年龄范围错误！")
		return
	}
	p.age = n
}

func (p *person) GetAge() int {
	return p.age
}

func (p *person) SetSalary(n float64) {
	if n <= 0 {
		fmt.Println("薪水有误！")
		return
	}
	p.salary = n
}

func (p *person) GetSalary() float64 {
	return p.salary
}
