package Proxy

import "fmt"

func ProxyTest()  {
	um:=UserManager{}
	users:=um.GetUserByDepId("00101")
	fmt.Printf("userId:%s userName:%s\n",users[0].GetUserId(),users[0].GetName())
	fmt.Printf("userId:%s userName:%s\n",users[1].GetUserId(),users[1].GetName())
	fmt.Printf("userId:%s userName:%s depId:%s Sex:%s\n",users[2].GetUserId(),users[2].GetName(),users[2].GetDepId(),users[2].GetSex())

	//保护代理
	o1:=Order{
		orderUser:"luke",
		productName:"T-shirt",
		orderNum:100,
	}
	var oApi OrderApi = &OrderProxy{
		OrderItem:o1,
	}
	//Jakc 尝试修改订单数
	oApi.SetOrderNum(200,"Jack")
	fmt.Printf("OrderNum:%d\n",oApi.GetOrderNum())
	//luke 本人尝试修改订单数
	oApi.SetOrderNum(200,"luke")
	fmt.Printf("OrderNum:%d\n",oApi.GetOrderNum())
}
