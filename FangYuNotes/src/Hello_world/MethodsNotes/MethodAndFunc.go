package MethodsNotes

import "fmt"

type Student struct {
	Name   string
	Gender string
	Age    int
	Id     string
	Score  float64
}

//func (student Student) String() string {
//	str := fmt.Sprintf("Name is %v\nGender is %v\nAge is %v\nId is %v\nScore is %v",
//			student.Name,student.Gender,student.Age,student.Id,student.Score)
//	return str
//}
func MethodAndFunc() {
	Tom := Student{
		Name:   "Tom",
		Gender: "TOM",
		Age:    20,
		Id:     "001",
		Score:  100,
	}
	fmt.Println(&Tom)
	Jerry := &Student{
		Name: "Jerry",
		Age:  19,
	}
	fmt.Println(Jerry)
	fmt.Println(&Jerry)
	fmt.Println(*Jerry)
}
