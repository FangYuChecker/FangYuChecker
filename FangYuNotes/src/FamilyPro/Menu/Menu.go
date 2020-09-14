package Menu2

import (
	"FamilyPro/Details"
	"FamilyPro/Init"
	"fmt"
)

func Menu() {
	key := ""
	for {
		fmt.Println("----------Family Note----------")
		fmt.Println("          1、收支明细          ")
		fmt.Println("          2、登记收入          ")
		fmt.Println("          3、登记支出          ")
		fmt.Println("          4、退出软件          ")
		fmt.Print("Please enter 1 - 4 :")
		fmt.Scanln(&key)

		switch key {
		case "1":
			Details.Details()
		case "2":
			Details.Income()
		case "3":
			Details.Expense()
		case "4":
			Init.Judge = false
		default:
			fmt.Println("Please enter a correct number !")
		}
		if !Init.Judge {
			break
		}
	}
	fmt.Println("Exit this system")
}
