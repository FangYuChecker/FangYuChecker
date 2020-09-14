package Abstract

import "fmt"

type Account struct {
	Name    string
	Pwd     string
	Age     int
	Balance float64
}

//存款
func (account *Account) Deposited(money float64, pwd string) {
	//比对密码
	if account.Pwd != pwd {
		fmt.Println("Password Wrong")
		return
	}
	if money <= 0 {
		fmt.Println("Money Wrong")
		return
	}
	account.Balance += money
	fmt.Println("Deposited Successful")
}

//取款
func (account *Account) Withdraw(money float64, pwd string) {
	//比对密码
	if account.Pwd != pwd {
		fmt.Println("Password Wrong")
		return
	}
	if money <= 0 || money > account.Balance {
		fmt.Println("Money Wrong")
		return
	}
	account.Balance -= money
	fmt.Println("Withdraw Successful")
}

//查询
func (account *Account) Query(pwd string) {
	if account.Pwd != pwd {
		fmt.Println("Password Wrong")
		return
	}
	fmt.Printf("Your Id is %v\nYour Balance is %v\n", account.Name, account.Balance)
}
func AbstractMethods() {
	Tom := Account{
		Name:    "Tom",
		Pwd:     "123",
		Age:     20,
		Balance: 300,
	}
	Tom.Query("123")
	Tom.Deposited(200, "123")
	Tom.Query("123")
	Tom.Withdraw(600, "123")
	Tom.Withdraw(200, "123")
	Tom.Query("123")
}
