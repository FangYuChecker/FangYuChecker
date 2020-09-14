package FileText

import (
	"fmt"
	"io/ioutil"
)

//读取文件的的内容并显示在终端（使用ioutil一次将整个文件读入到内存中）
//适用于文件不大的情况下
func FileApp2() {
	//打开文件

	file := "./src/FilerOperate/text.txt"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("file not found !", err)
	}

	//输出文件，看看文件是什么
	fmt.Printf("file =  %v\n", string(content))

	//没有显式的Open文件，因此也不需要显式的Close文件
	//因为，文件的Open和Close被封装到 ReadFile 函数内部
}
