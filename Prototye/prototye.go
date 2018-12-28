package Prototype

import "fmt"

type OrderApi interface {
	GetOrderProductNum() int
	SetOrderProductNum(num int)

	//克隆方法, @return 订单原型实例
	CloneOrder() OrderApi
}

//个人订单
type PersonalOrder struct {
	customerName string
	productId string
	orderProductNum int
}

func (po *PersonalOrder)GetOrderProductNum() int {
	return po.orderProductNum
}
func (po *PersonalOrder)SetOrderProductNum(num int)  {
	po.orderProductNum=num
}
func (po *PersonalOrder)GetCustomerName() string  {
	return po.customerName
}
func (po *PersonalOrder)SetCustomerName(name string)  {
	po.customerName=name
}
func (po *PersonalOrder)GetProductId() string  {
	return po.productId
}
func (po *PersonalOrder)SetProductId(prodId string)  {
	po.productId=prodId
}
func (po *PersonalOrder)String() string {
	return fmt.Sprintf("productID:%s customerName:%s orderProductNum:%d",po.productId,po.customerName,po.orderProductNum)
}

func (po *PersonalOrder)CloneOrder() OrderApi {
	newOrder:=PersonalOrder{}
	newOrder.SetCustomerName(po.customerName)
	newOrder.SetProductId(po.productId)
	newOrder.SetOrderProductNum(po.orderProductNum)

	return &newOrder
}
//企业订单
type EnterpriseOrder struct {
	enterpriseName string
	productId string
	orderProductNum int
}

func (eo *EnterpriseOrder)GetOrderProductNum() int {
	return eo.orderProductNum
}
func (eo *EnterpriseOrder)SetOrderProductNum(num int)  {
	eo.orderProductNum=num
}
func (eo *EnterpriseOrder)GetEnterpriseName() string  {
	return eo.enterpriseName
}
func (eo *EnterpriseOrder)SetEnterpriseName(name string)  {
	eo.enterpriseName=name
}
func (eo *EnterpriseOrder)GetProductId() string  {
	return eo.productId
}
func (eo *EnterpriseOrder)SetProductId(prodId string)  {
	eo.productId=prodId
}
func (eo *EnterpriseOrder)CloneOrder() OrderApi {
	newOrder:=EnterpriseOrder{}
	newOrder.SetEnterpriseName(eo.enterpriseName)
	newOrder.SetProductId(eo.productId)
	newOrder.SetOrderProductNum(eo.orderProductNum)

	return &newOrder
}
func (eo *EnterpriseOrder)String() string {
	return fmt.Sprintf("productID:%s enterpriseName:%s orderProductNum:%d",eo.productId,eo.enterpriseName,eo.orderProductNum)
}
