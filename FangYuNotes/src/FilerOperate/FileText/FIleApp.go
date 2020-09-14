package FileText

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func FileApp() {
	//打开文件
	file, err := os.Open("/Users/fangyuchecker/Desktop/GoProject/src/FilerOperate/text.txt")
	if err != nil {
		fmt.Println("file not found !", err)
	}

	//输出文件，看看文件是什么
	fmt.Printf("file =  %v\n", file)

	//当函数退出时候，要及时关闭文件
	//defer函数退出时候才会用到
	//要及时关闭文件句柄，否则会有内存泄漏
	defer file.Close()

	//创建一个reader，带缓冲的
	reader := bufio.NewReader(file)
	//循环读取文件内容
	for {
		//读到一个换行就结束一次
		str, error := reader.ReadString('\n')
		if error == io.EOF { // EOF表示文件的末尾
			break
		}
		fmt.Printf("%v", str)
	}

	fmt.Println("文件读取结束")
}
