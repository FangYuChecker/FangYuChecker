package InterfaceNote

import "fmt"

//声明一个接口
type USB interface {
	//声明了两个没有实现的方法
	Start()
	Stop()
}

type Phone struct {
}

//让结构体实现接口的方法
func (phone Phone) Start() {
	fmt.Println("Phone Start")
}
func (phone Phone) Stop() {
	fmt.Println("Phone Stop")
}

type Camera struct {
}

//让结构体实现接口的方法
func (camera Camera) Start() {
	fmt.Println("Camera Start")
}

func (camera Camera) Stop() {
	fmt.Println("Camera Stop")
}

//计算机
type Computer struct {
}

//编写工作方法,接收一个usb接口的变量
//只要是实现了usb接口，就是实现了USB接口u声明的所有方法
//着重提醒 是 声明的 所 有 方 法。
func (c Computer) Working(usb USB) {
	//通过USB接口变量调用start和stop
	usb.Start()
	usb.Stop()
}

func InterfaceNote() {
	c := Computer{}
	p := Phone{}
	a := Camera{}

	c.Working(p)
	c.Working(a)
}
