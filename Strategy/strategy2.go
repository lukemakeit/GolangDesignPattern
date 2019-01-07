package Strategy

import "fmt"

//支付工资的策略接口,公司有多种支付工资的算法
//如现金、银行卡、现金加股票、现金加期权、美元支付等
type PaymentStrategy interface {
	Pay(ctx *PaymentContext)
}
//人民币现金支付
type RMBCash struct {}
func (rc *RMBCash)Pay(ctx *PaymentContext)  {
	fmt.Printf("现在给 %s 人民币现金支付 %.2f元\n",ctx.UserName,ctx.Money)
}
//美元现金支付
type DollarCash struct {}
func (dc *DollarCash)Pay(ctx *PaymentContext)  {
	fmt.Printf("现在给 %s 美元现金支付 %.2f元\n",ctx.UserName,ctx.Money)
}
//银行卡支付
type BankCard struct {
	Account string
}
func (bc *BankCard)Pay(ctx *PaymentContext)  {
	fmt.Printf("现在给 %s 的账户%s转账 %.2f元\n",ctx.UserName,bc.Account,ctx.Money)
}

type PaymentContext struct {
	UserName string
	Money float64
	Strategy PaymentStrategy //支付工资的策略接口
}
func (pc *PaymentContext)PayNow()  {
	//使用客户希望的支付策略来支付工资
	pc.Strategy.Pay(pc)
}
