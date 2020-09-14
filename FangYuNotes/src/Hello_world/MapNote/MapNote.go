package MapNote

import "fmt"

//map is key-value var map_name map[keytype] valuetype
//slice,map,func不可以做为keytype
//通常为int string
//valuetype通常为int,string,map,struct
//var a map[string] string
//var a map[string] map[string]string
//声明不会分配内存，初始化需要用make，分配内存后才能赋值和使用

//结构体的定义
type Stu struct {
	Name    string
	Age     int
	Address string
}

func MapNote() {
	/*
		var a map[string]string
		//key不能重复，但是value可以重复
		//在这里fmt.Println(a) 为空值
		//golang 里面map是无序的结构
		a = make(map[string]string,10)
		a["no1"] = "song jiang"
		a["no2"] = "wu song"
		fmt.Println(a)
		a["no1"] = "wu yong"
		fmt.Println(a)
		a["no3"] = "wu yong"
		fmt.Println(a)
		//第二种命名方式
		cities := make(map[string]string)
		fmt.Println(cities)
		//第三种方式
		heroes := map[string]string{
			"hero_1" : "song jiang",
			"hero_2" : "lu jun yi",
		}
		fmt.Println(heroes )
	*/
	//map的增删查改

	/*var a map[string]string
	a = make(map[string]string,10)
	//增
	a["no1"] = "song jiang"
	a["no2"] = "wu song"
	a["no3"] = "wu yong"
	//改
	a["no1"] = "wu yong"
	*/
	//删除
	//delete(a,"no1")
	//fmt.Println(a)
	//删除map里面的所有key值，可以重新分配空间，或者遍历删除
	//map的查找
	/*
		val, ok := a["no2"]
		if ok {
			fmt.Println("no no1, value is",val)
		}else {
			fmt.Println("Yes")
		}
	*/
	//map的遍历只可以使用 for range
	//fmt.Println(a)
	//fmt.Println(len(a))
	//for s := range a {
	//	fmt.Println(s)
	//	fmt.Println(a[s])
	//}
	//演示map的切片使用
	//define a map's slice
	/*var monsters []map[string]string
	monsters = make([]map[string]string,1)
	//add a monster
	if monsters[0] == nil{
		monsters[0] = make(map[string]string)
		monsters[0]["name"] = "Tom"
		monsters[0]["age"] = "500"
	}
	fmt.Println(monsters)
	newMonster := map[string]string{
		"name" : "Jerry",
		"age" : "200",
	}
	monsters = append(monsters, newMonster)
	fmt.Println(monsters)
	*/
	//map的排序
	//a := make(map[int]int)
	//a[10] = 100
	//a[0] = 10
	//a[4] = 8
	//fmt.Println(a)
	//var keys []int
	//for i := range a {
	//	keys = append(keys, i)
	//}
	//sort.Ints(keys)
	//for i := range keys {
	//	fmt.Printf("a[%v] = %v\n",keys[i],a[keys[i]])
	//}
	Stus := make(map[string]Stu)
	stu1 := Stu{"Tom", 13, "tom"}
	stu2 := Stu{"Jerry", 14, "jerry"}
	Stus["1"] = stu1
	Stus["2"] = stu2
	for s := range Stus {
		fmt.Println(s)
		fmt.Println(Stus[s].Name, Stus[s].Age, Stus[s].Address)
	}
}

//
