package Channel

import "fmt"

//channel 的遍历和关闭
//channel关闭过后不能往里面写入数据，但是仍然可以读取数据
//

func ChannelList()  {
	intChan := make(chan int,10)
	intChan <- 100
	intChan <- 200
	intChan <- 300
	intChan <- 400
	//内建函数close，专门用来关闭channel
	close(intChan)
	//channel关闭过后才能正常遍历，如果没有关闭，会出现deadlock错误
	for i := range intChan {
		fmt.Println(i)
	}
}