package Memento

import "fmt"

//备忘录窄接口,外部通过这个窄接口是不知道内部数据情况的
type FlowMemento interface {}
//真正的备忘录对象,实现备忘录窄接口
//属性私有,不允许外部访问
type FlowMementoImpl struct {
	tempResult int //某个中间结果
	tempState string //某个中间状态
}

//模拟运行流程A
//相当于原发器对象,也就是需要备忘录保存数据的类
type FlowAMock struct {
	FlowName string //流程名称
	TempResult int //需要保存的,中间结果
	TempState string //需要保存的 中间状态
}
//运行流程的第一个阶段
func (f *FlowAMock)RunPhaseOne()  {
	f.TempResult=3
	f.TempState="PhaseOne"
}
//按照方案一运行流程后半部分
func (f *FlowAMock)Schema1()  {
	f.TempState=f.TempState+".Schema1"
	f.TempResult+=11
	fmt.Printf("%s:now run %d\n",f.TempState,f.TempResult)
}
//按照方案二运行流程后半部分
func (f *FlowAMock)Schema2()  {
	f.TempState=f.TempState+".Schema2"
	f.TempResult+=22
	fmt.Printf("%s:now run %d\n",f.TempState,f.TempResult)
}
//创建保存FlowAMock的备忘录对象,返回窄接口 外部不能访问到内部具体数据
func (f *FlowAMock)CreateMemento() FlowMemento{
	return &FlowMementoImpl{
		f.TempResult,
		f.TempState,
	}
}
//重新设置FlowAMock的状态,让其按照另一个访问继续运行
//也就是将备忘录的状态重置回去
func (f *FlowAMock)SetMemento(fm FlowMemento)  {
	fml:=fm.(*FlowMementoImpl)
	f.TempResult=fml.tempResult
	f.TempState=fml.tempState
}
//负责保存模拟流程A对象备忘录对象
type FlowMementoCareTaker struct {
	memnto  FlowMemento
}

func (fc *FlowMementoCareTaker)SaveMemento(fm FlowMemento)  {
	fc.memnto=fm
}
func (fc *FlowMementoCareTaker)RetriveMemento() FlowMemento  {
	return fc.memnto
}
