package Flyweight

import "fmt"

func FlyWeightTest()  {
	Login("张三")
	Login("李四")
	f1:=HasPermit("张三","人员列表","查看")
	f2:=HasPermit("李四","薪资数据","修改")
	f3:=HasPermit("张三","薪资数据","修改")
	fmt.Printf("f1==%v\n",f1)
	fmt.Printf("f2==%v\n",f2)
	fmt.Printf("f3==%v\n",f3)

	u1:=GetSecurityMgr()["张三"]
	u2:=GetSecurityMgr()["李四"]
	fmt.Println(u1)
	fmt.Println(u2)
}
