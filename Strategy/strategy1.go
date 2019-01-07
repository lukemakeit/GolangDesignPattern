package Strategy

import "fmt"

type CustormerStrategy interface {
	CalcPrice(goodsPrice float64) float64
}
//新客户或普通客户
type NormalCustomerStrategy struct {}
func (nc *NormalCustomerStrategy)CalcPrice(goodsPrice float64) float64 {
	fmt.Println("对于新客户或者普通客户,没有折扣")
	return goodsPrice
}
//老客户
type OldCustomerStrategy struct {}
func (oc *OldCustomerStrategy)CalcPrice(goodsPrice float64) float64 {
	fmt.Println("对于老客户,统一折扣5%")
	return goodsPrice*(1-0.05)
}
//大客户
type LargeCustomerStrategy struct {}
func (lc *LargeCustomerStrategy)CalcPrice(goodsPrice float64) float64  {
	fmt.Println("对大客户,统一折扣10%")
	return goodsPrice*(1-0.1)
}
//价格管理,向不同客户报不同价格
type Price struct {
	Strategy CustormerStrategy
}
func (p *Price)quote(goodsPrice float64) float64 {
	return p.Strategy.CalcPrice(goodsPrice)
}