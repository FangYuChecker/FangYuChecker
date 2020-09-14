package FileText

import (
	"fmt"
	"io/ioutil"
)

func TransferFile() {
	//将WriteFile文件导入到text文件中
	//首先将WriteFile文件读取到内存，将读取到的内容写入text
	filepathA := "./src/FilerOperate/WriteFile.txt"
	filepathB := "./src/FilerOperate/text.txt"

	data, err := ioutil.ReadFile(filepathA)
	if err != nil {
		fmt.Println("file not found !", err)
	}

	err = ioutil.WriteFile(filepathB, data, 0666)
	if err != nil {
		fmt.Println("Write File Error !")
	}
}
