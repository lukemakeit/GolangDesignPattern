package Strategy

import "fmt"

func StrategyTest()  {
	//不同客户的折扣策略
	cs:=&LargeCustomerStrategy{}
	os:=&OldCustomerStrategy{}

	//创建上下文
	ctx:=&Price{}
	ctx.Strategy=cs
	fmt.Printf("向客户报价:%.2f\n",ctx.quote(10000))

	ctx.Strategy=os
	fmt.Printf("向客户报价:%.2f\n",ctx.quote(10000))

	strategyRMB:=&RMBCash{}
	strategyDollar:=&DollarCash{}

	ctx1:=PaymentContext{
		"小李",
		5000,
		strategyRMB,
	}
	ctx1.PayNow()

	ctx2:=PaymentContext{
		"小张",
		8000,
		strategyDollar,
	}
	ctx2.PayNow()

	strategyCard:=&BankCard{
		"010998877665",
	}
	ctx3:=PaymentContext{
		"小王",
		9000,
		strategyCard,
	}
	ctx3.PayNow()
}
