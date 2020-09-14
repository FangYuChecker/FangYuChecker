package Channel

import (
	"fmt"
)

//使用select可以解决从管道取数据的堵塞问题

func App5()  {
	intChan := make(chan int,50)
	for i:=1 ;i<=50 ;i++ {
		//放入数据
		intChan <- i
		fmt.Println("writeData is ",i)
	}
	laber :
	for  {
		select {
			//注意：按这种方式，如果intChan最终都没关闭也不会一直阻塞在这里
			//没取到的话会自动到下一个case匹配
		case v:= <-intChan :
			fmt.Println(v)
		case v:= <-intChan:
			fmt.Println(v)
		default:
			fmt.Println("都取不到数字:）")
			break laber
		}
	}
}
