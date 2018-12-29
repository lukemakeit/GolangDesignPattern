package Facade

import "fmt"

//A模块接口
type AMoudleAPI interface {
	TestA() string
}
//A模块实现
type aMoudleImpl struct {}

func (am *aMoudleImpl)TestA() string {
	return "Yes,get a A moudle instance"
}
//A模块生成器
func NewAmoudleAPI() AMoudleAPI  {
	return &aMoudleImpl{}
}

//B模块接口
type BMoudleAPI interface {
	TestB() string
}
//B模块实现
type bMoudleImpl struct {}

func (bm *bMoudleImpl)TestB() string {
	return "Yes,get a B moudle instance"
}
//B模块生成器
func NewBmoudleAPI() BMoudleAPI  {
	return &bMoudleImpl{}
}

//系统整体API
type systemAPI interface {
	Test() string
}
//系统模块实现
type systemImpl struct {
	a AMoudleAPI
	b BMoudleAPI
}

func (si *systemImpl)Test() string {
	//系统统一调用各个模块功能
	aRet:=si.a.TestA()
	bRet:=si.b.TestB()
	return fmt.Sprintf("aRet:%s\nbRet:%s",aRet,bRet)
}
//系统模块生成器
func NewSystemMoudleAPI() systemAPI {
	return &systemImpl{
		a:NewAmoudleAPI(),
		b:NewBmoudleAPI(),
	}
}

