package Details

import (
	"FamilyPro/Init"
	"fmt"
)

func Income() {
	var income float64
	for {
		fmt.Println("Please enter your income : ")
		fmt.Scanln(&income)
		if income < 0 {
			fmt.Println("Data Error , Redo !")
		} else {
			break
		}
	}
	fmt.Println("您的本次收入：", income)
	Init.User1.Balance += income
	fmt.Println("您的收入后的总金额为：", Init.User1.Balance)
}
