package Visitor

import "fmt"

//顾客类型接口
type Customer interface {
	Accept(v Visitor)
}
//访问者接口
type Visitor interface {
	Visit(c Customer)
}
//企业用户
type EnterpriseCustormer struct {
	CusomerId string `json:"customer_id"`
	Name string `json:"name"`	//客户名称
	LinkMan string `json:"link_man"` //联系人
	LinkTelePhone string `json:"link_telephone"` //联系人电话
	RegisterAddress string `json:"register_address"` //注册地址
}
func (ec *EnterpriseCustormer)Accept(v Visitor)  {
	v.Visit(ec)
}

//个人用户
type PersonalCustormer struct {
	CusomerId string `json:"customer_id"`
	Name string `json:"name"`	//客户名称
	Telephone string `json:"telephone"` //客户电话
	Age int `json:"age"` //年龄
}
func (pc *PersonalCustormer)Accept(v Visitor)  {
	v.Visit(pc)
}

//服务请求功能
type ServiceRequestVisitor struct {}
func (srv *ServiceRequestVisitor)Visit(c Customer)  {
	switch cc:=c.(type) {
	case *EnterpriseCustormer:
		fmt.Printf("%s企业提出服务请求\n",cc.Name)
	case *PersonalCustormer:
		fmt.Printf("客户%s提出服务请求\n",cc.Name)
	}
}

//客户偏好分析
type PredctionAnalazeVisitor struct {}
func (pav *PredctionAnalazeVisitor)Visit(c Customer)  {
	switch cc:=c.(type) {
	case *EnterpriseCustormer:
		fmt.Printf("现在对企业客户%s进行产品偏好分析\n",cc.Name)
	case *PersonalCustormer:
		fmt.Printf("现在对个人客户%s进行产品偏好分析\n",cc.Name)
	}
}
//ObjectStructure的实现
type CustomerCol struct {
	customers []Customer
}

func (c *CustomerCol) Add(customer Customer) {
	c.customers = append(c.customers, customer)
}

func (c *CustomerCol) Accept(visitor Visitor) {
	for _, customer := range c.customers {
		customer.Accept(visitor)
	}
}