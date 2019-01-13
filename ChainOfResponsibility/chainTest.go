package ChainOfResponsibility

import "fmt"

func ChainTest()  {
	h1:=&GeneralManager{}
	h2:=&DepManager{}
	h3:=&ProjectManager{}
	h3.Successor=h2
	h2.Successor=h1

	ret1:=h3.HandleFeeRequest("张三",1200)
	fmt.Println(ret1)
	ret2:=h3.HandleFeeRequest("小李",1200)
	fmt.Println(ret2)

	//业务处理对象
	ebo:=&GoodsSaleEbo{}
	saleM:=SaleModel{
		Goods:"张学友怀旧经典",
		SaleNum:100,
	}
	ebo.Sale("小李","张三",saleM)
}
