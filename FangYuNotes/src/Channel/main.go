package main

import "Channel/Channel"

//要求计算1-200的各个数的阶乘，放到map里面，要求用到goroutine
//出现资源争夺问题
//channel的本质就是一个队列
//数据先进先出
//线程安全
//channel是有类型的
//channel是引用类型，而且要初始化
//var (
//	myMap = make(map[int]int)
//	//u声明一个全局的互斥锁
//	//lock 是一个全局的互斥锁
//	//synchornized
//	//Mutes 互斥
//	lock sync.Mutex
//)
//
//func test(n int)  {
//	var res int
//	res = 1
//	for i:=1; i<=n; i++{
//		res *= i
//	}
//	lock.Lock()
//	myMap[n] = res
//	lock.Unlock()
//}
func main()  {
	Channel.App6()
	//for i:=1; i<=20; i++{
	//	go test(i)
	//}
	//time.Sleep(time.Second*5)
	//lock.Lock()
	//for i := range myMap {
	//	fmt.Printf("map[%v]=%v\n",i,myMap[i])
	//}
	//lock.Unlock()

	//channel可以是只读或只写，默认情况下都是双向
	//只写
	//var intChan chan<-int
	//intChan = make(chan int,3 )
	//intChan <- 20
	//num := <-intChan //会报错
	//只读
	//var intChan<-chan int
	//intChan <- 20//会报错
	//
}
