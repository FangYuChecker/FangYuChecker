package FileText

import (
	"flag"
	"fmt"
)

//Flag包处理命令行的各种参数
//flag包实现了命令行参数的解析
func FlagCmd() {
	//定义几个参数，用于接收命令行的参数
	var user string
	var pwd string
	var host string
	var port int

	//&user 就是接收用户命令行中的输入的-u 后面的参数值
	//"u" ，就是-u指定参数
	//""就是默认值
	//"用户名，默认为空" ，说明
	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&pwd, "p", "", "密码，默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名，默认为localhost")
	flag.IntVar(&port, "port", 3306, "默认3306")
	//转换，必须调用该方法
	flag.Parse()
	fmt.Println(user, pwd, host, port)

}
