package Proxy

import "fmt"

/*
* 订单系统,一旦订单被创建,则只有订单创建人才可以修改订单中的数据,其他不能修改
 */
 //订单对象接口
 type OrderApi interface {
 	GetProductName() string
 	SetProductName(productName string,user string)
 	GetOrderNum() int
 	SetOrderNum(orderNum int,user string)
 	GetOrderUser() string
 	SetOrderUser(orderUser string,user string)
 }
 //订单对象
 type Order struct {
 	productName string
 	orderNum int
 	orderUser string
 }
func (o *Order)GetProductName() string {
	return o.productName
}
 //Order中所有的修改都么有设置权限,proxy中将做权限检查
func (o *Order)SetProductName(productName string,user string)  {
	o.productName=productName
}
func (o *Order)GetOrderNum() int{
	return o.orderNum
}
func (o *Order)SetOrderNum(orderNum int,user string)  {
	o.orderNum=orderNum
}
func (o *Order)GetOrderUser() string{
	return o.orderUser
}
func (o *Order)SetOrderUser(orderUser string,user string)  {
	o.orderUser=orderUser
}

type OrderProxy struct {
	OrderItem Order
}
func (op *OrderProxy)GetProductName() string {
	return op.OrderItem.GetProductName()
}
//OrderProxy中将做权限检查
func (op *OrderProxy)SetProductName(productName string,user string)  {
	if op.OrderItem.GetOrderUser()== user{
		op.OrderItem.SetOrderUser(productName,user)
	}else{
		fmt.Printf("Sorry %s,只有订单属主才有权限修改订单数据\n",user)
	}
}
func (op *OrderProxy)GetOrderNum() int{
	return op.OrderItem.GetOrderNum()
}
func (op *OrderProxy)SetOrderNum(orderNum int,user string)  {
	if op.OrderItem.GetOrderUser() == user{
		op.OrderItem.SetOrderNum(orderNum,user)
	}else{
		fmt.Printf("Sorry %s,只有订单属主才有权限修改订单数据\n",user)
	}
}
func (op *OrderProxy)GetOrderUser() string{
	return op.OrderItem.GetOrderUser()
}
func (op *OrderProxy)SetOrderUser(orderUser string,user string)  {
	if op.OrderItem.GetOrderUser() == user{
		op.OrderItem.SetOrderUser(orderUser,user)
	}else{
		fmt.Printf("Sorry %s,只有订单属主才有权限修改订单数据\n",user)
	}
}
