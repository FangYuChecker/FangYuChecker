package Menu

import (
	"CustomerInformationManagerSystem/Init"
	"CustomerInformationManagerSystem/OperateCustomerInformation"
	"fmt"
)

type customerView struct {
	key             string
	loop            bool
	customerService *OperateCustomerInformation.CustomerService
}

func (this *customerView) list() {
	//获取当前所有的客户信息
	customers := this.customerService.List()
	fmt.Println("----------------------客户列表----------------------")
	fmt.Println()
	for i := range customers {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println()
	fmt.Println("--------------------客户列表完成---------------------")
}

func (this *customerView) add() {
	var Name string
	var Gender string
	var Age int
	var Phone string
	var Email string
	fmt.Printf("请输入姓名：")
	fmt.Scanln(&Name)
	fmt.Printf("请输入性别：")
	fmt.Scanln(&Gender)
	fmt.Printf("请输入年龄：")
	fmt.Scanln(&Age)
	fmt.Printf("请输入手机号：")
	fmt.Scanln(&Phone)
	fmt.Printf("请输邮箱地址：")
	fmt.Scanln(&Email)
	customer := Init.NewCustomer2(Name, Gender, Age, Phone, Email)
	if this.customerService.AddCustomer(customer) {
		fmt.Println("添加完成")
	} else {
		fmt.Println("添加失败")
	}
}

func (this *customerView) delete() {
	id := -1
	fmt.Printf("请输入删除客户的id(-1退出) ：")
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	this.customerService.DeleteCustomer(id)
}
func (this *customerView) showMenu() {
	for {
		fmt.Println("----------Customer System----------")
		fmt.Println("            1、添加客户          ")
		fmt.Println("            2、修改客户          ")
		fmt.Println("            3、删除客户          ")
		fmt.Println("            4、客户列表          ")
		fmt.Println("            5、退    出          ")
		fmt.Print("Please enter 1 - 5 :")
		fmt.Scanln(&this.key)

		switch this.key {
		case "1":
			this.add()
		case "2":
			OperateCustomerInformation.ChangeCustomer()
		case "3":
			this.delete()
		case "4":
			this.list()
		case "5":
			this.loop = false
		default:
			fmt.Println("Please enter a correct number !")
		}
		if !this.loop {
			break
		}
	}
	fmt.Println("Exit this system")
}

func Menu() {
	customerView := customerView{
		key:  "",
		loop: true,
	}
	customerView.customerService = OperateCustomerInformation.NewCustomerService()
	customerView.showMenu()
}
