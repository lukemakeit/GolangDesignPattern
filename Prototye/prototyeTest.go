package Prototype

import "fmt"

func SaveOrder(api OrderApi)  {
	for api.GetOrderProductNum()>1000{
		nOrder:=api.CloneOrder()
		nOrder.SetOrderProductNum(1000)
		api.SetOrderProductNum(api.GetOrderProductNum()-1000)
		fmt.Printf("拆分成订单:%s\n",nOrder)
	}
	fmt.Printf("订单是:%s\n",api)
}

func PrototypeTest()  {
	pOrder:=PersonalOrder{}
	pOrder.SetProductId("P0001")
	pOrder.SetCustomerName("张三")
	pOrder.SetOrderProductNum(2925)

	SaveOrder(&pOrder)
}