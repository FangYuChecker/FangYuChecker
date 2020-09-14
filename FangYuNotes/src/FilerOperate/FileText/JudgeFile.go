package FileText

import (
	"fmt"
	"os"
)

//golang判断文件或文件夹是否存在使用os.Stat()函数返回的错误值进行判断
//如果返回的是nil，说明文件或文件夹存在
//如果返回的错误类型使用os.IsNotExist()判断为true，说明文件或文件夹不存在
//如果返回其他类型，说明不确实是否存在

func pathExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}
func JudgeFile() {
	if pathExist("./src") {
		fmt.Println("File is ok !")
	}
}
