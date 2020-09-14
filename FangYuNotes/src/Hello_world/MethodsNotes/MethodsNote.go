package MethodsNotes

import "fmt"

//Golang 的方法是作用在指定的数据类型上的，
//和指定的数据类型绑定，因此自定义类型都可以有方法
type Person struct {
	Name string
	Age  int
}

func (Tom Person) Test() {
	fmt.Printf("%v is a good boy\n", Tom.Name)
}

func (Tom Person) Calculate() {
	Sum := 0
	for i := 1; i <= Tom.Age; i++ {
		Sum += i
	}
	fmt.Printf("Sum is %v\n", Sum)
}

func (Tom Person) GetSum(n1 int, n2 int) int {
	Sum := 0
	for i := n1; i <= n2; i++ {
		Sum += n1
	}
	return Sum
}

func (Tom *Person) GetSum2(n1 int, n2 int) int {
	Sum := 0
	for i := n1; i <= n2; i++ {
		Sum += n1
	}
	return Sum
}

type Integer int

func (i Integer) Print() {
	fmt.Println(i)
}

//如果一个类型实现了String（）这个方法，fmt.Println会默认调用这个方法进行输出
func (person Person) String() string {
	str := fmt.Sprintln("Tom is a Cat", person.Age)
	return str
}
func MethodsNotes() {
	var i Integer = 10
	i.Print()

	Tom := Person{"Tom", 10}
	fmt.Println(&Tom)

	//Test() 不能直接使用，只能通过Person来调用，而不能使用。
	Tom.Test()
	Tom.Calculate()
	fmt.Println(Tom.GetSum(10, 20))
	fmt.Println(Tom.GetSum2(10, 20))
}
