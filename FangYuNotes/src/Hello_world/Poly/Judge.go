package Poly

import "fmt"

func TypeJudge(items ...interface{}) {
	for i := range items {
		switch items[i].(type) {
		case bool:
			fmt.Println("Is a bool", items[i])
		case int, int8, int16, int32, int64:
			fmt.Println("Is a int", items[i])
		default:
			fmt.Println("Unknown")
		}
	}
}
func JudgeType() {
	a := 32
	b := 33.3
	c := true
	TypeJudge(a, b, c)
}
