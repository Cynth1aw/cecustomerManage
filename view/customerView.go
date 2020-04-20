package main

import (
	"fmt"
	"laurencecustomerManage/service"
	"laurencecustomerManage/model"
)

type customerView struct {
	option string
	loop bool
	customerService *service.CustomerService
}
func (cv *customerView) list() {
	customer := cv.customerService.List()
	fmt.Println("---------------------------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t邮箱\t电话\n ")
	for i := 0; i < len(customer); i++ {
		fmt.Println(customer[i].GetInfo())
	}
	fmt.Println("---------------------------------------")
}
func(cv *customerView) add() {
	fmt.Println("-------------添加客户-----------------")
	name := ""
	fmt.Print("姓名：")
	fmt.Scanln(&name)
	for {
		if name != "" {
			break
		}
		fmt.Print("姓名：")
		fmt.Scanln(&name)
	}
	
	gender := ""
	fmt.Print("性别：")
	fmt.Scanln(&gender)
	for {
		if gender != "" {
			break
		}
		fmt.Print("性别：")
		fmt.Scanln(&gender)
	}
	age := 0
	fmt.Print("年龄：")
	fmt.Scanln(&age)
	for {
		if age != 0 {
			break
		}
		fmt.Print("年龄：")
		fmt.Scanln(&age)
	}
	email := ""
	fmt.Print("邮箱：")
	fmt.Scanln(&email)
	for {
		if email != "" {
			break
		}
		fmt.Print("邮箱：")
		fmt.Scanln(&email)
	}
	phone := ""
	fmt.Print("电话：")
	fmt.Scanln(&phone)
	for {
		if phone != "" {
			break
		}
		fmt.Print("电话：")
		fmt.Scanln(&phone)
	}
	//构建Customer结构体
	customer := model.NewCustomerNoId(name,gender,age,email,phone)
	cv.customerService.Add(customer)
	fmt.Println("--------------------------------------")
}
func (cv *customerView) delete() {
	fmt.Print("请输入要删除的ID：")
	id := -1
	fmt.Scanln(&id)
	fmt.Println("确认删除? y/n：")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Print("输入有误，请重新输入：")
	}
	if choice == "y" && cv.customerService.Delete(id) {
		fmt.Println("删除成功")
	} else {
		fmt.Println("删除失败")
	}
}
func (cv *customerView) exit() {
	fmt.Print("您确定要退出吗? y/n：")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Print("输入有误，请重新输入：")
	}
	if choice == "y" {
		cv.loop = false
	}
}
func (cv *customerView) modifile() {
	fmt.Print("请输入要编辑的用户ID: ")
	id := -1
	fmt.Scanln(&id)
	indx, customer := cv.customerService.FindByIdModifile(id)
	fmt.Println(customer[0])
	if customer != nil {
		fmt.Println("--------编辑更新客户信息---------")
		fmt.Printf("姓名(%v)：",customer[0].Name)
		name := ""
		fmt.Scanln(&name)
		fmt.Printf("性别(%v)：",customer[0].Gender)
		gender := ""
		fmt.Scanln(&gender)
		fmt.Printf("年龄(%v)：",customer[0].Age)
		age := 0
		fmt.Scanln(&age)
		fmt.Printf("电话(%v)：",customer[0].Phone)
		phone := ""
		fmt.Scanln(&phone)
		fmt.Printf("邮箱%v)：",customer[0].Email)
		email := ""
		fmt.Scanln(&email)
		upCustomer := model.NewCustomer(id,name,gender,age,phone,email)
		if cv.customerService.Updat(indx,upCustomer) {
			fmt.Println("更新成功")
		} else {
			fmt.Println("更新成功")
		}
	}
}
func (cv *customerView) mainMenu() {
	for {
		fmt.Println("----------家庭收支记账明细------------")
		fmt.Println("             1 添加客户")
		fmt.Println("             2 修改客户")
		fmt.Println("             3 删除客户")
		fmt.Println("             4 客户列表")
		fmt.Println("             5 退    出")
		fmt.Println("             请选择(1-5)")
		fmt.Println("--------------------------------------")
		fmt.Print("请选择您要操作的选项：")
		fmt.Scanln(&cv.option)
		switch cv.option {
		case "1":
			cv.add()
		case "2":
			cv.modifile()
		case "3":
			cv.delete()
		case "4":
			cv.list()
		case "5":
			cv.exit()
		default:
			fmt.Print("输入有误，请重新输入：")
		}
		if !cv.loop {
			fmt.Println("欢迎下次光临，Goodbay!!!!!!!!!!!!")
			break
		}
	}
}
func main() {
	var cv customerView = customerView{
		option: "",
		loop: true,
	}
	cv.customerService = service.NewCustomerService()
	cv.mainMenu()
}
