package ArrayNote

import (
	"fmt"
	"math/rand"
	"time"
)

func ArrayNotes() {

	var intArray [3]int
	//为了每次生成的随机数不同，需要给一个seed value
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(intArray); i++ {
		intArray[i] = rand.Intn(100) //小于等于100
	}
	fmt.Println(intArray)
	//for range
	//for index/_,value := range Array
	//heroes := [...]string{"song jiang"}
	/*
		var score [6]float64
		for i:=1 ;i <=len(score) ;i++{
			fmt.Printf("Please enter a value : ")
			fmt.Scanln(&score[i-1])
		}
		fmt.Println(score)
	*/
}
