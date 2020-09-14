package ArrayNote

import (
	"fmt"
)

func SliceNotes() {

	var a []int = []int{1, 2, 3, 4}
	var slice = make([]int, 1)
	fmt.Println(slice)
	copy(slice, a)
	//只会拷贝第一个元素
	fmt.Println(slice)

	//演示切片slice的使用make
	/*
		var slice []float64 =make([]float64,5,10)
		fmt.Println(slice)
		fmt.Println(cap(slice))
		fmt.Println(len(slice))
		slice = append(slice, 4,2,3,4,5,1,3,3,3,3,3,3,3,4,2,3)
		fmt.Println(slice)
		fmt.Println(len(slice))
		fmt.Println(cap(slice))
		//for i := range slice {
		//	fmt.Println(slice[i])
		//	fmt.Println(i)
		//}
		slice[2]=1
		sort.Float64s(slice)
		fmt.Println(slice)
	*/

	//var slice []string = []string{"Tom","Marry"}
	//声明定义一个切片
	//slice := intArr[:]
	//slice := intArr[1:3]
	//slice 就是切片名字
	//intArr[1:3] 表示slice引用到intArr这个数组
	//引用intArr的起始下标为1，到3，但是不包含3
	//slice包含 address,len,cap
	/*
		var intArr [5]int = [...]int{10,20,30,40,50}
		slice := intArr[1:3]
		fmt.Println(intArr)
		fmt.Println(slice)
		fmt.Println("切片的容量是：",cap(slice)) //容量是动态变化的
	*/
}
