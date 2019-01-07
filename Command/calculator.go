package Command

import "fmt"

//操作运算的接口,其实这里也不用提供接口,难到操作运算加法有多种实现,减法有多种实现么?所以感觉提供一个类即可
//不过对于执行开机 关机命令的主板来说可以提供一个接口,因为主板可是有很多供应商的
type OperateApi interface {
	GetResult() int
	SetResult(result int)
	Add(num int)
	Substract(num int)
}
//真正实现加减法运算(receiver)
type Operation struct {
	result int
}
func (o *Operation)GetResult() int  {
	return o.result
}
func (o *Operation)SetResult(result int)  {
	o.result=result
}
func (o *Operation)Add(num int)  {
	o.result=o.result+num
}
func (o *Operation)Substract(num int)  {
	o.result=o.result-num
}

//命令接口,有多个命令(加 减 乘 除),所以统一一个interface是很好的
type Command interface {
	Execute()
	Undo()
}

//加法命令
type AddCommand struct {
	Operator OperateApi //真正执行加法操作的receiver
	OpeNum int //加法操作数
}
func (add *AddCommand)Execute()  {
	add.Operator.Add(add.OpeNum) //receiver来执行加法操作
}
func (add *AddCommand)Undo()  {
	add.Operator.Substract(add.OpeNum)//加法 对应的 就是减法
}
//减法命令
type SubstractCommand struct {
	Operator OperateApi //真正执行减法操作的receriver
	OpeNum int //减法操作数
}
func (sub *SubstractCommand)Execute()  {
	sub.Operator.Substract(sub.OpeNum) //receiver来执行减法操作
}
func (sub *SubstractCommand)Undo()  {
	sub.Operator.Add(sub.OpeNum) //减法 对应的 就是加法
}

//计算器实现(Invoker)
type Calculator struct {
	AddCmd Command
	SubCmd Command
	UndoCmds []Command
	RedoCmds []Command
}
func (c *Calculator)AddPressed()  {
	c.AddCmd.Execute()
	//把操作记录记录到Undo历史中
	c.UndoCmds=append(c.UndoCmds,c.AddCmd)
}
func (c *Calculator)SubstractPressed()  {
	c.SubCmd.Execute()
	//把操作记录记录到Undo历史中
	c.UndoCmds=append(c.UndoCmds,c.SubCmd)
}
func (c *Calculator)UndoPressed()  {
	if len(c.UndoCmds)>0 {
		//取出最后一个命令来执行撤销
		cmd:=c.UndoCmds[len(c.UndoCmds)-1]
		cmd.Undo()
		//恢复功能,这里把命令记录到恢复历史中
		c.RedoCmds=append(c.RedoCmds,cmd)
		//把最后一个命令删除
		c.UndoCmds=c.UndoCmds[:len(c.UndoCmds)-1]
	}else{
		fmt.Println("很抱歉，没有可撤销的命令了")
	}
}
func (c *Calculator)RedoPressed()  {
	if len(c.RedoCmds)>0 {
		//取出最后一个命令来执行重做
		cmd:=c.RedoCmds[len(c.RedoCmds)-1]
		cmd.Execute()
		//这里把命令记录到可撤销的历史中
		c.UndoCmds=append(c.UndoCmds,cmd)
		//把最后一个命令删除
		c.RedoCmds=c.RedoCmds[:len(c.RedoCmds)-1]
	}else{
		fmt.Println("很抱歉，没有可恢复的命令了")
	}
}
