package ArrayNote

func StringAndSlice() {
	//修改字符串，转为arr数组，修改，再为字符串，但是无法修改中文
	/*
		str := "hello@atguigu"
		fmt.Println(str)
		arr :=[]byte(str)
		fmt.Println(arr)
		arr[0] = 'z'
		str = string(arr)
		fmt.Println(str)
	*/
	//中文部分可以用[]rune进行修改，之后再转回去
	/*
		str := "hello@atguigu"
		fmt.Println(str)
		arr :=[]rune (str)
		fmt.Println(arr)
		arr[0] = '北'
		str = string(arr)
		fmt.Println(str)
	*/
	//字符串和切片的部分获取
	/*
		str := "hello@atguigu"
		//切片获取atguigu
		slice := str[6:]
		fmt.Println(slice)
	*/
}
