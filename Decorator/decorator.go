package Decorator

import (
	"fmt"
)

//计算奖金的组件接口
// @param user被计算奖金的人员
// @param begin 计算奖金的开始时间
// @param end 计算奖金的结束时间
type Component interface {
	CalcPrize(user string,begin string,end string) float64
}

//计算奖金的类
type ConcreteComponent struct {}
func (cc *ConcreteComponent)CalcPrize(user string,begin string,end string ) float64  {
	return 0
}

//装饰器接口Decorator就没有了,我们来实现具体的装饰器类吧
//装饰器对象,当月业务奖金
type MonthPrizeDecorator struct {
	c Component
}
func (md *MonthPrizeDecorator)CalcPrize(user string,begin string,end string) float64 {
	money:=md.c.CalcPrize(user,begin,end)
	prize,ok:=GetTempDB()[user]
	if ok==false {
		prize=0.0
	}
	fmt.Printf("%s当月业务奖金:%.2f\n",user,prize*0.03)
	return money+prize*0.03
}
//当月业务奖金构造方法
func NewMonthPrizeDeco(c Component) Component {
	return &MonthPrizeDecorator{
		c:c,
	}
}

//装饰器对象,累计奖金
type SumPrizeDecorator struct {
	c Component
}
func (sd *SumPrizeDecorator)CalcPrize(user string,begin string,end string) float64 {
	money:=sd.c.CalcPrize(user,begin,end)
	//假定大家的累计业务额都市1000000元
	prize:=1000000*0.01
	fmt.Printf("%s累计奖金:%.2f\n",user,prize)
	return money+prize
}
//累计奖金构造函数
func NewSumPrizeDeco(c Component) Component {
	return &SumPrizeDecorator{
		c:c,
	}
}

//装饰对象,团队业务奖金,只有leader才有
type GroupPrizeDecorator struct {
	c Component
}
func (gd *GroupPrizeDecorator)CalcPrize(user string,begin string,end string) float64 {
	money:=gd.c.CalcPrize(user,begin,end)
	//计算当月团队业务奖金,先计算出团队总业务额,然后乘以1%
	group:=0.0
	for _,v:=range GetTempDB(){
		group=group+v
	}
	fmt.Printf("%s当月团队业务奖金:%.2f\n",user,group*0.01)
	return money+group*0.01
}
func NewGroupPrizeDeco(c Component) Component {
	return &GroupPrizeDecorator{
		c:c,
	}
}