package ArrayNote

import "fmt"

func DoubleArr() {
	//四行六列的数组
	var arr [4][6]int
	for i := range arr {
		for i2 := range arr[i] {
			fmt.Printf("%v ", arr[i][i2])
		}
		fmt.Println()
	}
	fmt.Println(arr)
}
