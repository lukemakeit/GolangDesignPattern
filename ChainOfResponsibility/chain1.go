package ChainOfResponsibility

import "fmt"

//职责对象接口
type Handler interface {
	//处理聚餐费用的申请
	//@param user 申请人
	//@param fee 申请的钱数
	//@return 成功或失败的通知
	HandleFeeRequest(user string,fee float64) string
}
//项目经理处理
type ProjectManager struct {
	Successor Handler //下一个处理请求对象
}
func (pm *ProjectManager)HandleFeeRequest(user string,fee float64) string  {
	var str string
	if fee<500{
		//为了测试简单,只同意小李的
		if user== "小李"{
			str=fmt.Sprintf("项目经理同意%s %.2f元的聚餐费用",user,fee)
		}else{
			str=fmt.Sprintf("项目经理不同意%s %.2f元的聚餐费用",user,fee)
		}
		return str
	}else{
		if pm.Successor!=nil{
			return pm.Successor.HandleFeeRequest(user,fee)
		}
	}
	return str
}
//部门经理处理
type DepManager struct {
	Successor Handler //下一个处理请求对象
}
func (dm *DepManager)HandleFeeRequest(user string,fee float64) string  {
	var str string
	if fee<1000{
		//为了测试简单,只同意小李的
		if user== "小李"{
			str=fmt.Sprintf("部门经理同意%s %.2f元的聚餐费用",user,fee)
		}else{
			str=fmt.Sprintf("部门经理不同意%s %.2f元的聚餐费用",user,fee)
		}
		return str
	}else{
		if dm.Successor!=nil{
			return dm.Successor.HandleFeeRequest(user,fee)
		}
	}
	return str
}
//总经理处理
type GeneralManager struct {
	Successor Handler //下一个处理请求对象
}
func (gm *GeneralManager)HandleFeeRequest(user string,fee float64) string  {
	var str string
	if fee>=1000{
		//为了测试简单,只同意小李的
		if user== "小李"{
			str=fmt.Sprintf("总经理同意%s %.2f元的聚餐费用",user,fee)
		}else{
			str=fmt.Sprintf("总经理不同意%s %.2f元的聚餐费用",user,fee)
		}
		return str
	}else{
		if gm.Successor!=nil{
			return gm.Successor.HandleFeeRequest(user,fee)
		}
	}
	return str
}
