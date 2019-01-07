package Decorator

import "fmt"

func DecoratorTest()  {
	//普通员工奖金计算
	//基本奖金
	c1:=&ConcreteComponent{}
	//组合当月业务奖金
	md:=NewMonthPrizeDeco(c1)
	//组合累计奖金
	sd:=NewSumPrizeDeco(md)
	zsMoney:=sd.CalcPrize("张三","","")
	fmt.Printf("======张三应得奖金额:%.2f\n",zsMoney)
	lsMoney:=sd.CalcPrize("李四","","")
	fmt.Printf("======李四应得奖金额:%.2f\n",lsMoney)

	//leader,还需要加上团队奖金计算
	gd:=NewGroupPrizeDeco(sd)
	wwMoney:=gd.CalcPrize("王五","","")
	fmt.Printf("======王五应得奖金额:%.2f\n",wwMoney)
}
