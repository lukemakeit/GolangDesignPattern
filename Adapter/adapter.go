package Adapter

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type LogModel struct {
	LogId string `json:"logId"` //日志id
	Operator string `json:"operator"` //操作人
	OperateTime string `json:"operateTime"` //操作时间
	LogContent string `json:"logContent"` //日志内容
}
//日志文件操作接口
type LogFileOperateApi interface {
	// 读取日志文件,获取日志列表
	ReadLogFile()([]LogModel)
	//写日志文件,把日志列表写入日志文件中
	WriteLogFile([]LogModel)
}
//实现日志文件操作
type LogFileOperate struct {
	LogFilePathName string `json:"logFilePathName"` //日志文件路径和文件名称
}

func (lfo *LogFileOperate)ReadLogFile()([]LogModel) {
	LogList:= make([]LogModel,0,3)
	//这里就不真的去读取数据了,模拟几条读取到的数据即可
	LogList=append(LogList,LogModel{
		LogId:"1",
		Operator:"张三",
		OperateTime:"2018-12-28 10:00:00",
		LogContent:"restart serever:1.1.1.1",
	})
	LogList=append(LogList,LogModel{
		LogId:"2",
		Operator:"李四",
		OperateTime:"2018-12-29 11:00:00",
		LogContent:"restart serever:2.2.2.2",
	})
	LogList=append(LogList,LogModel{
		LogId:"3",
		Operator:"王五",
		OperateTime:"2018-12-30 11:00:00",
		LogContent:"restart serever:3.3.3.3",
	})
	return LogList
}
func (lfo *LogFileOperate)WriteLogFile(logList []LogModel)  {
	//这里就直输出到终端吧,就不输出到文件了
	for _,item:=range logList{
		fmt.Printf("logId:%s\tOperator:%s\tOperateTime:%sLogContent:%s\n",item.LogId,item.Operator,item.OperateTime,item.LogContent)
	}
}

//日志DB操作接口
type LogDbOpeateApi interface {
	//新增日志
	CreateLog(log LogModel)
	//修改日志
	UpdateLog(log LogModel)
	//删除日志
	RemoveLog(log LogModel)
	//获取所有日志
	GetAllLog()([]LogModel)
}
//日志DB操作
type LogDbOperate struct {
	DBConn *gorm.DB  `json:"-"` //DB连接
	TableName string `json:"tableName"`
}
func (lto *LogDbOperate)CreateLog(log LogModel)  {
	//简单打印一下SQL语句即可
	fmt.Printf("insert into %s(logId,operator,operateTime,logContent)values('%s','%s','%s','%s');\n",lto.TableName,log.LogId,log.Operator,log.OperateTime,log.LogContent)
}
func (lto *LogDbOperate)UpdateLog(log LogModel)  {
	//简单打印一下SQL语句即可
	fmt.Printf("update %s set logId='%s',operator='%s',operateTime='%s',logContent='%s' where logId='%s';\n",lto.TableName,log.LogId,log.Operator,log.OperateTime,log.LogContent,log.LogId)
}
func (lto *LogDbOperate)RemoveLog(log LogModel)  {
	//简单打印一下SQL语句即可
	fmt.Printf("delete from %s  where logId='%s';\n",lto.TableName,log.LogId)
}
func (lto *LogDbOperate)GetAllLog()([]LogModel)  {
	//简单打印一下SQL语句即可
	fmt.Printf("select * from %s;\n",lto.TableName)
	LogList:= make([]LogModel,0,3)
	//这里就不真的去读取数据了,模拟几条读取到的数据即可
	LogList=append(LogList,LogModel{
		LogId:"1",
		Operator:"张三",
		OperateTime:"2018-12-28 10:00:00",
		LogContent:"restart serever:1.1.1.1",
	})
	LogList=append(LogList,LogModel{
		LogId:"2",
		Operator:"李四",
		OperateTime:"2018-12-29 11:00:00",
		LogContent:"restart serever:2.2.2.2",
	})
	LogList=append(LogList,LogModel{
		LogId:"3",
		Operator:"王五",
		OperateTime:"2018-12-30 11:00:00",
		LogContent:"restart serever:3.3.3.3",
	})
	return LogList
}
//适配器
type LogFileDBAdapter struct {
	FileAdaptee LogFileOperateApi
}
func (lfd *LogFileDBAdapter)CreateLog(log LogModel)  {
	//调用LogFileOperateApi 接口的WriteLogFile()方法来实现新建日志
	lfd.FileAdaptee.WriteLogFile([]LogModel{log})
}
func (lfd *LogFileDBAdapter)UpdateLog(log LogModel)  {
	//先把已有的日志全部读取出来
	logList:=lfd.FileAdaptee.ReadLogFile()
	//遍历已有日志
	for idx,item:=range logList{
		if item.LogId == log.LogId{
			//修改原值
			logList[idx].LogId = log.LogId
			logList[idx].Operator = log.Operator
			logList[idx].OperateTime = log.OperateTime
			logList[idx].LogContent = log.LogContent
			break
		}
	}
	//调用LogFileOperateApi接口的WriteLogFile()方法把修改后的日志写回去,成本有点高
	//先清空日志文件,在全部写回
	lfd.FileAdaptee.WriteLogFile(logList)
}
func (lfd *LogFileDBAdapter)RemoveLog(log LogModel)  {
	//先把已有的日志全部读取出来
	logList:=lfd.FileAdaptee.ReadLogFile()
	logRet:=[]LogModel{}
	//遍历已有日志
	for idx,item:=range logList{
		if item.LogId == log.LogId{
			//loglist只有一个元素,idx=0
			if idx==0 && len(logList)==0{
				break
			}else if idx ==0{
				//logList不止一个元素,idx=0
				logRet=logList[1:]
				break
			}else if len(logList)==idx+1{
				//logLit不止一个元素,idx指向logList的最后一个元素
				logRet=logList[:idx-1]
			}else{
				//logList不止一个元素,idx不等于0,同时idx也不指向logList的最后一个元素
				logRet=logList[:idx]
				logRet=append(logRet,logList[idx+1:]...)
			}
			break
		}
	}
	//调用LogFileOperateApi接口的WriteLogFile()方法把修改后的日志写回去,成本有点高
	//先清空日志文件,在全部写回
	lfd.FileAdaptee.WriteLogFile(logList)
}
func (lfd *LogFileDBAdapter)GetAllLog()([]LogModel) {
	//调用LogFileOperateApi 接口的ReadLogFile()方法来实现
	return lfd.FileAdaptee.ReadLogFile()
}
