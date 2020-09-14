package List

import (
	"encoding/json"
	"fmt"
)

type Cat struct {
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Salary   float64 `json:"salary"`
	Birthday string  `json:"birthday"`
}

func testStruct() {
	cat := Cat{
		Name:     "Tom",
		Age:      20,
		Salary:   999.9,
		Birthday: "19970721",
	}
	//将cat序列化
	data, err := json.Marshal(&cat)
	if err != nil {
		fmt.Println("序列化失败：", err)
	}
	//输出序列化后的结果
	fmt.Printf("Cat序列化为：%v\n", string(data))
}

func testMap() {
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "Jerry"
	a["age"] = 30
	a["address"] = "XHGY"
	//map是引用方式传递，不需要&
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Println("序列化失败：", err)
	}
	//输出序列化后的结果
	fmt.Printf("a map序列化为：%v\n", string(data))
}

func testSlice() {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "Jerry"
	m1["age"] = 30
	m1["address"] = [3]string{"XHGY", "YZGY", "THGY"}
	slice = append(slice, m1)
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Println("序列化失败：", err)
	}
	//输出序列化后的结果
	fmt.Printf("slice序列化为：%v\n", string(data))
}
func StructList() {
	//演示将结构体，map，切片序列化
	testStruct()
	testMap()
	testSlice()
}
