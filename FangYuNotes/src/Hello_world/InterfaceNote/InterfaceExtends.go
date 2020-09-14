package InterfaceNote

import "fmt"

type A interface {
	testA()
}
type B interface {
	testB()
}
type C interface {
	A
	B
	testC()
}
type Student struct {
}

func (stu Student) testA() {
	fmt.Println("A")
}
func (stu Student) testB() {
	fmt.Println("B")
}
func (stu Student) testC() {
	fmt.Println("C")
}
func InterfaceExtends() {
	var stu Student
	a := C(stu)
	stu.testA()
	a.testB()
	a.testC()
}
