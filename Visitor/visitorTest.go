package Visitor

func VisitorTest()  {
	col:=&CustomerCol{}
	col.Add(&EnterpriseCustormer{
		Name:"ABC集团",
	})
	col.Add(&EnterpriseCustormer{
		Name:"CDE公司",
	})
	col.Add(&PersonalCustormer{
		Name:"张三",
	})
	//服务请求
	srv:=&ServiceRequestVisitor{}
	//偏好分析
	pav:=&PredctionAnalazeVisitor{}
	col.Accept(srv)
	col.Accept(pav)

}
