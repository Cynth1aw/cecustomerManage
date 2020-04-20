package service
import (
	_"fmt"
	"laurencecustomerManage/model"
)

type CustomerService struct {
	customerNum int
	customers []model.Customer
}
func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(customerService.customerNum,"lauren7ce","男",28,"lauren7ce@outlook.com","13538239520")
	customerService.customers = append(customerService.customers,*customer)
	return customerService
}
func (cs *CustomerService) Delete(id int) bool {
	index := cs.FindById(id)
	if index == -1 {
		return false
	}
	//这里等下要来理解理解
	cs.customers = append(cs.customers[:index],cs.customers[index + 1:]...)
	return true
}
func (cs *CustomerService) FindById(id int) int {
	index := -1
	for i, _ := range cs.customers {
		if cs.customers[i].Id == id {
			index = i
			break
		}
	}
	return index
}
func (cs *CustomerService) FindByIdModifile(id int) (key int, customer []model.Customer) {
	index := -1
	for i, _ := range cs.customers {
		if cs.customers[i].Id == id {
			index = i
			break
		}
	}
	if index != -1 {
		customer := cs.customers[index:index+1]
		return index, customer
	} else {
		return 0, nil
	}
}
func (this *CustomerService) Updat(index int,customer *model.Customer) bool {
	this.customers[index] = *customer
	return true
}
func (cs *CustomerService) List() []model.Customer {
	return cs.customers
}
func (cs *CustomerService) Add(customer *model.Customer) bool {
	cs.customerNum += 1 
	customer.Id = cs.customerNum
	cs.customers = append(cs.customers,*customer)
	return true
}