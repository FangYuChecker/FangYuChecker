package StringNote

//func len(str)
//解决中文乱码
//for ... range
// []rune(str)
//字符串转数字 n,err := strconv.Atoi("12")
//整数转字符串 str =strconv.Itoa(123)
//10进制转换 str = strconv.FormatInt(i int64, base int) //string
//查找子串  strings.Contains("seafood","foo") //true,false
//返回字符串s 中有几个不重复的sep子串 func strings.Count(s, sep string) //int
//不区分字母大小写的比较 strings.EqualFold("abc","Abc")，区分大小写直接 “Abc” == “abc”就可以
//返回子串在字符串第一次出现的index值，否则返回-1 strings.Index("Nlt_abc","abc") //int
//返回子串在字符串最后一次出现的index值，否则返回-1 strings.LastIndex("Nlt_abc","abc") //int
//将指定的字串替换成另一个字串 strings.Replace("go go hello","gp","go language",n)
//指定字符串拆分出多个字符数组 strings.Split("hello,world,ok",",")
//字符串大小写转换 strings.ToLower("str") // strings.ToUpper("str")
//去除两边的空格 或字符 strings.TrimSpace(str) strings.Trim("!!!!hello###","!#")
//strings.TrimLeft
//strings.TrimRight
//判断字符串是否以指定的开头strings.HasPrefix(str1,str2) //bool
//判断字符串是否以指定的结尾strings.HasSuffix(str1,str2) //bool
func StringNotes() {
	/*
		str := "go lang GO LANG"
		str = strings.ToLower(str)
		fmt.Printf("%v",str)
		str = strings.ToUpper(str)
		fmt.Printf("%v",str)
	*/
	/*str := strings.Split("hello,world,ok",",")
	fmt.Printf("%v",str)
	*/
	/*
		str := strings.Replace("go go hello" , "go","北京",-1)
		fmt.Printf("str = %v\n",str)
	*/

	//查找子串  strings.Contains("seafood","foo") //true,false
	/*
	   b := strings.Contains("seafood","foo")
	   fmt.Printf("b = %v\n", b)
	*/
	//字符串转数字 n,err := strconv.Atoi("12")
	/*
	   n,err := strconv.Atoi("123")
	   if err != nil {
	       fmt.Println("转换错误",err)
	   }else{
	       fmt.Println("转成的结果是：",n)
	   }
	*/
	//遍历含有中文的字符串
	/*
	   str2 := []rune(str)
	   for i := 0 ; i<len(str2) ;i++{
	       fmt.Printf("Type is %c\n",str2[i])
	   }
	   str1 := ""
	   for index, val := range str1 {
	       //不得不用的变量，但是可以忽略他
	       _ = index
	       fmt.Printf(" val = %c \n", val )
	   }
	*/
}
