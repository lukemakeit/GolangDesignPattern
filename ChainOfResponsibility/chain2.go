package ChainOfResponsibility

import "fmt"

//权限检查=>通用数据校验=>数据逻辑校验=>业务逻辑实现
//简单封装销售单的数据
type SaleModel struct {
	Goods string `json:"goods"` //商品名称
	SaleNum int `json:"sale_num"` //销售数量
}
func (sm *SaleModel)String() string  {
	return fmt.Sprintf("商品名称=%s,销售数量=%d",sm.Goods,sm.SaleNum)
}
//职责对象接口
type SaleHandler interface {
	//处理保存销售信息的请求
	//@param user 操作人员
	//@param customer 客户
	//@param saleModel 销售数据
	//@return 是否处理成功
	Sale(user string,customer string,saleModle SaleModel) bool
}
//安全检查职责对象
type SaleSecurityCheck struct {
	//下一个职责对象
	Successor SaleHandler
}
func (sc *SaleSecurityCheck)Sale(user string,customer string,saleModel SaleModel) bool  {
	//权限检查,简单只让操作员小李通过
	if user == "小李" {
		fmt.Printf("恭喜你,%s.保存销售数据权限检查通过\n",user)
		return sc.Successor.Sale(user,customer,saleModel)
	}else{
		fmt.Printf("对不起,%s.你没有保存销售数据的权限\n",user)
	}
	return false
}
//数据通用检查职责对象
type SaleDataCheck struct {
	//下一个职责对象
	Successor SaleHandler
}
func (sdc *SaleDataCheck)Sale(user string,customer string,saleModel SaleModel) bool  {
	//数据通用检查
	if len(user)==0 {
		fmt.Println("操作申请人不能为空")
		return false
	}
	if len(customer)==0{
		fmt.Println("客户名称不能为空")
		return false
	}
	if len(saleModel.Goods)==0{
		fmt.Println("销售商品名称不能为空")
		return false
	}
	if saleModel.SaleNum<=0 {
		fmt.Printf("销售商品数量不能小于0,%s=>%d\n",saleModel.Goods,saleModel.SaleNum)
		return false
	}
	return sdc.Successor.Sale(user,customer,saleModel)
}
//数据逻辑检查职责对象
type SaleLogicCheck struct {
	//下一个职责对象
	Successor SaleHandler
}
func (slc *SaleLogicCheck)Sale(user string,customer string,saleModel SaleModel) bool  {
	//数据逻辑检查,比如检查产品在数据库中是否存在,产品ID是否唯一
	//这里就不做示例了,直接过
	return slc.Successor.Sale(user,customer,saleModel)
}
//真正处理销售业务的职责对象
type SaleMgr struct {
	//下一个职责对象
	Successor SaleHandler
}
func (sm *SaleMgr)Sale(user string,customer string,saleModel SaleModel) bool  {
	//数据逻辑检查,比如检查产品在数据库中是否存在,产品ID是否唯一
	//这里就不做示例了,直接过
	fmt.Printf("%s保存了%s购买'%s'的销售数据",user,customer,saleModel)
	return true
}

//组织职责链,进行业务处理
type GoodsSaleEbo struct {
	//保存销售数据
}
func (gs *GoodsSaleEbo)Sale(user string,customer string,saleModel SaleModel) bool {
	//1.权限检查
	//2.通用数据检查
	//3.数据逻辑校验
	//4.真正业务处理
	ssc:=&SaleSecurityCheck{}
	sdc:=&SaleDataCheck{}
	slc:=&SaleLogicCheck{}
	sm:=&SaleMgr{}
	ssc.Successor=sdc
	sdc.Successor=slc
	slc.Successor=sm

	return ssc.Sale(user,customer,saleModel)
}