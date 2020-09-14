package OperateCustomerInformation

import (
	"CustomerInformationManagerSystem/Init"
)

type CustomerService struct {
	customers   []Init.Customer
	customerNum int
}

func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := Init.NewCustomer(1, "Tom", "female", 10,
		"18166310390", "18616310390@163.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

//返回客户切片
func (this *CustomerService) List() []Init.Customer {
	return this.customers
}

func (this *CustomerService) AddCustomer(customer Init.Customer) bool {
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers, customer)
	return true
}

func ChangeCustomer() {

}

func (this *CustomerService) DeleteCustomer(id int) bool {
	index := this.FindById(id)
	if index == -1 {
		return false
	}
	//如何从切片中删除一个元素
	this.customers = append(this.customers[:index], this.customers[index+1:]...)
	return true
}

func (this *CustomerService) FindById(id int) int {
	index := -1
	for i := 0; i < len(this.customers); i++ {
		if this.customers[i].Id == id {
			index = i
		}
	}
	return index
}
