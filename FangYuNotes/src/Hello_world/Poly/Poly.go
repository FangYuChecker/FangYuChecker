package Poly

import "fmt"

//变量具有多种形态，是面向对象的第三大特征
//go中，多态是通过接口实现的
//可以通过统一的接口调用不同的实现，这样接口就实现不同的形态

//接口体现多态的两种形式
//多态参数
//多态数组

//声明一个接口
type USB interface {
	//声明了两个没有实现的方法
	Start()
	Stop()
}

type Phone struct {
	name string
}

//让结构体实现接口的方法
func (phone Phone) Start() {
	fmt.Println("Phone Start")
}
func (phone Phone) Stop() {
	fmt.Println("Phone Stop")
}
func (phone Phone) Call() {
	fmt.Println("Phone Call")
}

type Camera struct {
	name string
}

//让结构体实现接口的方法
func (camera Camera) Start() {
	fmt.Println("Camera Start")
}

func (camera Camera) Stop() {
	fmt.Println("Camera Stop")
}

type Computer struct {
}

func (computer Computer) working(usb USB) {
	usb.Start()
	//如果usb是指向一个phone结构体，还需要调用call方法
	//类型断言
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	} else {
		fmt.Println("NOT A PHONE")
	}
	usb.Stop()
}
func Poly() {
	var usbArr [3]USB
	//fmt.Println(usbArr)
	usbArr[0] = Phone{"xnmi"}
	usbArr[1] = Phone{"huawei"}
	usbArr[2] = Camera{"yamaxun"}
	//fmt.Println(usbArr)
	var computer Computer
	for i := range usbArr {
		computer.working(usbArr[i])
		fmt.Println()
	}
}
