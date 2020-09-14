package Details

import (
	"FamilyPro/Init"
	"fmt"
)

func Expense() {
	var expense float64
	for {
		fmt.Println("Please enter your income : ")
		fmt.Scanln(&expense)
		if expense < 0 || expense > Init.User1.Balance {
			fmt.Println("Data Error , Redo !")
		} else {
			break
		}
	}
	fmt.Println("您的本次支出：", expense)
	Init.User1.Balance -= expense
	fmt.Println("您的支出后的总金额为：", Init.User1.Balance)
}
