package main

import "FilerOperate/FileText"

//文件在程序中是以流的形式呈现的
//获取文件的信息
//type Count struct {
//	En int
//	No int
//	Sp int
//	Other int
//}
func main() {
	FileText.WriteFile()
	//counter := Count{
	//	En:    0,
	//	No:    0,
	//	Sp:    0,
	//	Other: 0,
	//}
	//filepath := "./src/FilerOperate/WriteFile.txt"
	//file,err := os.OpenFile(filepath,os.O_RDWR | os.O_APPEND, 0666)
	//if err != nil {
	//	fmt.Println("File Open Error")
	//	return
	//}
	////及时关闭
	//defer file.Close()
	//
	//reader := bufio.NewReader(file)
	//for  {
	//	str,err := reader.ReadString('\n')
	//	if err == io.EOF{
	//		break
	//	}
	//	for i := range str {
	//		if str[i] >='0' && str[i] <='9' {
	//			counter.No++
	//		} else if (str[i] >='a' && str[i] <='z')||(str[i] >='A' && str[i] <='Z') {
	//			counter.En++
	//		} else if str[i] == ' ' || str[i] == '\t'{
	//			counter.Sp++
	//		}else {
	//			counter.Other++
	//		}
	//	}
	//}
	//fmt.Println(counter)
}
