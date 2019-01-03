package Flyweight

var testDB *TestDB

type TestDB struct {
	colDB []string
}
func (td *TestDB)GetUserAndPrivs() []string  {
	if td.colDB==nil{
		td.colDB=make([]string,0,4)
		td.colDB=append(td.colDB,"张三,人员列表,查看")
		td.colDB=append(td.colDB,"李四,人员列表,查看")
		td.colDB=append(td.colDB,"李四,薪资数据,查看")
		td.colDB=append(td.colDB,"李四,薪资数据,修改")
	}
	return td.colDB
}
func init()  {
	testDB=&TestDB{}
}