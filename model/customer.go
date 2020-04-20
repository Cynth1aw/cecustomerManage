package model
import "fmt"
type Customer struct {
	Id int 
	Name string
	Gender string
	Age int
	Email string
	Phone string
}
func NewCustomer(id int,name string,gander string,age int, email string,phone string) *Customer {
	return &Customer{
		Id: id,
		Name: name,
		Gender: gander,
		Age: age,
		Email: email,
		Phone: phone,
	}
}
func NewCustomerNoId(name string,gander string,age int, email string,phone string) *Customer {
	return &Customer{
		Name: name,
		Gender: gander,
		Age: age,
		Email: email,
		Phone: phone,
	}
}
func (customer *Customer) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v",customer.Id,customer.Name,customer.Gender,customer.Age,customer.Email,customer.Phone)
	return info
}