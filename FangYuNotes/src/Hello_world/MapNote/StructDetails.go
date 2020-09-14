package MapNote

import (
	"encoding/json"
	"fmt"
)

//struct 上每一个字段可以写一个tag，tag的可以通过反射机制获取，常见的就是序列化和反序列化
type Monster struct {
	Name  string
	Age   int
	Skill string
}
type Points struct {
	x int
	y int
}

type Rect struct {
	LeftP, RightP *Points
}

func Refershes(Point *Rect, x1 int, y1 int, x2 int, y2 int) {
	Point.LeftP.x = x1
	Point.LeftP.y = y1
	Point.RightP.x = x2
	Point.RightP.y = y2
}
func StructDetails() {
	//creat a monster
	monster := Monster{"Tom", 10, "Catch"}
	//将变量序列化为json格式字串
	jsonMonster, error := json.Marshal(monster)
	if error != nil {
		fmt.Println("Error", error)
	} else {
		fmt.Println(string(jsonMonster))
	}
	//var r1 Rect
	// r1 := Rect{&Points{1,3},&Points{3,4}}
	////r1 = new(Rect)
	//Refershes(&r1,4,2,3,4)
	//fmt.Println(&r1,r1,*r1.RightP)
}
