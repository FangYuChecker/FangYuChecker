package Channel

import (
	"fmt"
	"math"
	"time"
)

func putNum(intChan chan int)  {
	for i:=1 ;i<=8000 ;i++{
		intChan <- i
	}
	close(intChan)
}

func primeNum(intChan chan int,priChan chan int,exitChan chan bool)  {
	for{
		num,ok :=<- intChan
		if !ok{
			break
		}
		if isPrime(num) {
			priChan <- num
		}
	}
	exitChan <- true
	//fmt.Println("一个线程取不到数字，退出")
}

func isPrime(n int) bool {
	if n<=3{
		return n>1
	}
	if n%2==0 ||n%3==0{
		return false
	}
	for i:=5 ;i<int(math.Sqrt(float64(n)))+1 ;i++{
		if n%i==0 || n%(i+2)==0{
			return false
		}
	}
	return true
}

func App4()  {
	intChan := make(chan int ,1000)
	priChan := make(chan int ,8000)
	exitChan := make(chan bool,4)

	start := time.Now().Unix()
	//fmt.Println(start)
	go putNum(intChan)

	for i:=0 ;i<4 ;i++ {
		go primeNum(intChan,priChan,exitChan)
	}

	go func() {
		for i:=0;i<4 ;i++{
			<-exitChan
		}

		end := time.Now().Unix()
		fmt.Println("使用协程时间耗时：",end-start)
		fmt.Println("end = ",end)
		close(priChan)

	}()

	//遍历priChan
	for  {
		_,ok := <-priChan
		if !ok{
			break
		}
		//fmt.Println("素数：",v)
	}
	fmt.Println(start)
	fmt.Println("主线程退出")
}
