package Memento

func MementoTest()  {
	mock:=FlowAMock{
		FlowName:"TestFlow",
	}
	//运行流程第一个阶段
	mock.RunPhaseOne()

	fc:=FlowMementoCareTaker{}
	//保存数据
	fc.SaveMemento(mock.CreateMemento())
	//方案一运行后面步骤
	mock.Schema1()
	//恢复数据
	mock.SetMemento(fc.RetriveMemento())
	//方案二运行后面步骤
	mock.Schema2()
}

