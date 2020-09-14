package FileText

import (
	"fmt"
	"os"
)

//获取命令行的各种参数os.Args
func GetCmd() {
	fmt.Println("命令行的参数有：", len(os.Args))
	//变量os切片
	for i := range os.Args {
		fmt.Printf("Args[%v] = %v\n", i, os.Args[i])
	}
}
