package FileText

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func WriteFile() {
	filepath := "./src/FilerOperate/WriteFile.txt"
	//os.O_WRONLY os.O_RDONLY 只读只写文件打开 //os.O_RDWR
	//os.O_APPEND 打开文件并且加到文件末尾
	//os.O_CREATE 打开文件，不存在就创建
	//os.O_TRUNC 打开已有文件，并清空
	//file,err := os.OpenFile(filepath,os.O_WRONLY | os.O_CREATE, 0666)
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("File Open Error")
		return
	}
	//及时关闭
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}
	//准备写入五句话,\n表示换行，有一些编辑器看的是\r
	str := "Hello_World2\r\n"
	//写入时，使用带缓存的*Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}

	//因为Writer是带缓存的，因此调用writerstring方法
	//其实是先写入到缓存的，所以需要调用flush这个方法，将缓存的数据真正写入到文件中
	//否则文件中没有数据
	writer.Flush()
}
