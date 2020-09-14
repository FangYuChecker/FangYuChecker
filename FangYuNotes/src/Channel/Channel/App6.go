package Channel

import (
	"fmt"
	"time"
)

func sayH()  {
	for i:=0 ;i<10 ;i++{
		time.Sleep(time.Second)
		fmt.Println("Say Hello !")
	}
}
func test()  {
	defer func() {
		if err:= recover(); err !=nil{
			 fmt.Println("test error !" )
		}
	}()
	var map1 map[int]string
	map1[0] = "Tom"
}
func App6()  {
	go sayH()
	go test()

	for i:=0 ;i<10 ;i++{
		fmt.Println("main ok ",i)
		time.Sleep(time.Second)
	}
}
