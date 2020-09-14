package MapNote

import "fmt"

type Cat struct {
	Name  string
	Age   int
	Color string
	Ptr   *int
	Slice []int
	Map1  map[string]string
}

//如果结构体的字段类型 为 指针，map，slice的零值都是nil，即没有分配空间
//若想使用，则需要用make
func StructNotes() {
	var cat1 Cat
	cat1.Age = 15
	cat1.Name = "Tom"
	cat1.Color = "White"
	fmt.Println(cat1)
	if cat1.Ptr == nil && cat1.Slice == nil && cat1.Map1 == nil {
		fmt.Println("OK")
	}
	//指针使用需要new一下
	cat1.Ptr = new(int)
	*cat1.Ptr = 10
	cat1.Slice = make([]int, 10)
	cat1.Map1 = make(map[string]string)
	cat1.Slice[0] = 10
	cat1.Map1["Color"] = "White"
	fmt.Println(*cat1.Ptr)

	//指针结构体
	cat2 := new(Cat)
	//标准赋值方式
	(*cat2).Age = 10
	//go设计者在底层会对下面语言进行处理
	//会给cat2进行取值运算，也不会报错，但是内部的指针要记得new一下
	cat2.Name = "Jerry"
	fmt.Println(*cat2)
}
