package Channel

import "fmt"

type Car struct {
	name string
	age int
}
func ChannelTest()  {
	//演示管道的使用
	var intChan chan int
	intChan = make(chan int,3)
	//看看intChan是什么
	fmt.Println(intChan)
	fmt.Println(&intChan)
	//向管道写入数据
	intChan <- 10
	num := 211
	intChan <- num
	//看看管道的长度和cap（容量）
	fmt.Println(len(intChan), cap(intChan))
	//从管道中读取数据
	var num2 int
	num2 = <- intChan
	fmt.Println(num2)
	fmt.Println(len(intChan), cap(intChan))
	//在没有使用协程情况下，管道数据如果全部取出，会报错deadlock

	var intChan2 chan interface{}
	intChan2 = make(chan interface{},10)
	cat1 := Car{
		name: "Tom",
		age:  20,
	}
	intChan2 <- cat1
	intChan2 <- 10
	intChan2 <- "Name"
	//取出
	cat1s := <- intChan2
	fmt.Println(cat1s)
	//类型断言
	a := cat1s.(Car)
	fmt.Println(a.name)
}
