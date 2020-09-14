package FileText

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//拷贝文件，io包提供

//编写一个函数，获取源文件路径和目标文件路径
func getPath(filepath string, pathSrc string) (written int64, err error) {
	srcFile, err1 := os.Open(pathSrc)
	if err1 != nil {
		fmt.Println("Open File Error Src! ")
	}
	defer srcFile.Close()
	//通过src句柄获取reader
	reader := bufio.NewReader(srcFile)

	disFile, err2 := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err2 != nil {
		fmt.Println("File Open Error Dis !")
		return
	}
	defer disFile.Close()
	writer := bufio.NewWriter(disFile)

	return io.Copy(writer, reader)
}

func CopyFile() {
	disPath := "./src/FilerOperate/test.txt"
	srcPath := "./src/FilerOperate/text.txt"
	_, err := getPath(disPath, srcPath)
	if err == nil {
		fmt.Println("Copy Complete !")
	} else {
		fmt.Println("Copy Error !")
	}
}
