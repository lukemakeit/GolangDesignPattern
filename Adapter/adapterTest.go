package Adapter

import "fmt"

func AdapterTest()  {
	lm1:=LogModel{
		LogId:"100",
		Operator:"admin",
		OperateTime:"2015-10-11 10:03:11",
		LogContent:"这是一个测试",
	}
	logList:=[]LogModel{}
	logList=append(logList,lm1)

	//日志文件操作接口
	var logFileApi LogFileOperateApi
	logFileApi=&LogFileOperate{
		LogFilePathName:"log001.log",
	}
	logFileApi.WriteLogFile(logList)

	ll:=logFileApi.ReadLogFile()
	for _,item:=range ll{
		fmt.Println(item)
	}

	//日志DB操作接口
	var logDBApi LogDbOpeateApi
	logDBApi=&LogDbOperate{
		DBConn:nil,
		TableName:"logTable",
	}
	logDBApi.CreateLog(lm1)

	//日志DB操作接口,但是我们实际用适配器进行赋值
	var logDBApi2 LogDbOpeateApi
	logDBApi2=&LogFileDBAdapter{
		FileAdaptee:logFileApi,
	}
	logDBApi2.CreateLog(lm1)
}
