package FileText

import (
	"fmt"
	"os"
)

func FileOpenCLose() {
	//打开文件
	file, err := os.Open("/Users/fangyuchecker/Desktop/GoProject/src/FilerOperate/text.txt")
	if err != nil {
		fmt.Println("file not found !", err)
	}

	//输出文件，看看文件是什么
	fmt.Printf("file =  %v\n", file)

	//关闭文件
	err = file.Close()
	if err != nil {
		fmt.Println("Close file error !")
	}
}
