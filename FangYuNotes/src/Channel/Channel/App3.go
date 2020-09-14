package Channel

import "fmt"

func writeData(intChan chan int)  {
	for i:=1 ;i<=50 ;i++ {
		//放入数据
		intChan <- i
		fmt.Println("writeData is ",i)
	}
	//关闭
	close(intChan)
}

func readData(intChan chan int,exitChan chan bool)  {
	for  {
		v,ok := <- intChan
		if !ok{
			break
		}
		fmt.Println("readData is ",v)
	}
	exitChan <- true
	close(exitChan)
}
func App3() {
	intChan := make(chan int,50)
	exitChan := make(chan bool,1)

	go writeData(intChan)
	go readData(intChan,exitChan)

	for  {
		_,ok := <- exitChan
		if !ok {
			break
		}
	}

}
