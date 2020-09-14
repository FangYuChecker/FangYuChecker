package List

import (
	"encoding/json"
	"fmt"
)

//将json字符串反序列化为struct，map和slice
//保证反序列化后的数据类型和数据化之前的数据类型一致
func unmarshalStruct() {
	str := "{\"name\":\"Tom\",\"age\":20,\"salary\":999.9,\"birthday\":\"19970721\"}"
	//通常str是通过网络传输获取的
	//定义一个Cat实例
	var cat Cat
	//只有引用&cat传递才能改变函数外面的值
	err := json.Unmarshal([]byte(str), &cat)
	if err != nil {
		fmt.Printf("Unmarshal Error  = %v \n!", err)
	}
	fmt.Println("反序列化后， cat = ", cat)
}
func unmarshalMap() {
	str := "{\"name\":\"Tom\",\"age\":20,\"salary\":999.9,\"birthday\":\"19970721\"}"
	//通常str是通过网络传输获取的
	var a map[string]interface{}
	//反序列化底层会自动make空间
	//只有引用&cat传递才能改变函数外面的值
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Printf("Unmarshal Error  = %v \n!", err)
	}
	fmt.Println("反序列化后， a = ", a)
}
func unmarshalSlice() {
	str := "[{\"address\":[\"XHGY\",\"YZGY\",\"THGY\"],\"age\":30,\"name\":\"Jerry\"}]"
	//通常str是通过网络传输获取的
	var slice []map[string]interface{}
	//反序列化底层会自动make空间
	//只有引用&cat传递才能改变函数外面的值
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Printf("Unmarshal Error  = %v \n!", err)
	}
	fmt.Println("反序列化后， slice = ", slice)
}
func ReverseList() {
	unmarshalStruct()
	unmarshalMap()
	unmarshalSlice()

}
