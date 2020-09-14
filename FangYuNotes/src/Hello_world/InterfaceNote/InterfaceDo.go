package InterfaceNote

import (
	"fmt"
	"math/rand"
	"sort"
)

type Stus struct {
	Name string
	Age  int
}

//声明结构体类型切片
type StuSlice []Stus

//实现slice接口
func (Stu StuSlice) Len() int {
	return len(Stu)
}

//less决定你使用什么标准进行排序
//1 按年龄从小到大排序
func (Stu StuSlice) Less(i int, j int) bool {
	return Stu[i].Age > Stu[j].Age
}

func (Stu StuSlice) Swap(i int, j int) {
	//temp := Stu[i]
	//Stu[i] = Stu[j]
	//Stu[j] = temp
	//下面一句等于上面三句
	Stu[i], Stu[j] = Stu[j], Stu[i]
}
func InterfaceDo() {
	//define a slice
	Slice := []int{10, 2, 4, 5, 3, 9, 7}
	//sort slice
	sort.Ints(Slice)
	fmt.Println(Slice)

	//对结构体切片进行排序
	//使用系统提供的方法
	//test
	var stu StuSlice
	for i := 0; i < 10; i++ {
		n := Stus{
			Name: fmt.Sprintf("Stus is %d", rand.Intn(100)),
			Age:  i,
		}
		stu = append(stu, n)
	}
	for i := range stu {
		fmt.Println(stu[i])
	}
	sort.Sort(stu)
	fmt.Println("After Sort......")
	for i := range stu {
		fmt.Println(stu[i])
	}
}
