package Command

import "fmt"

func CalculaterTest()  {
	//receiver
	opera:=&Operation{
		result:0,
	}
	//创建命令对象并组装命令
	addCmd:=&AddCommand{
		Operator:opera,
		OpeNum:5,
	}
	subCmd:=&SubstractCommand{
		Operator:opera,
		OpeNum:3,
	}

	//把命令设置到持有者也就是计算器中
	calcu:=&Calculator{
		AddCmd:addCmd,
		SubCmd:subCmd,
		UndoCmds: make([]Command, 0),
		RedoCmds: make([]Command,0),
	}

	//模拟按下按钮
	calcu.AddPressed()
	fmt.Printf("一次加法运算后结果:%d\n",opera.GetResult())
	calcu.SubstractPressed()
	fmt.Printf("一次减法运算后的结果:%d\n",opera.GetResult())

	//测试撤销
	//subCmd.OpeNum=8
	calcu.UndoPressed()
	fmt.Printf("撤销一次运算后的结果:%d\n",opera.GetResult())
	calcu.UndoPressed()
	fmt.Printf("再次撤销运算后的结果:%d\n",opera.GetResult())

	//测试恢复
	calcu.RedoPressed()
	fmt.Printf("恢复一次运算后的结果:%d\n",opera.GetResult())
}
//计算器只能加5，只能减2。如果我将AddCommand中的OpeNum从5改成8的话，那我的撤销操作就会出问题，因为保存undoCmds redoCmds中保存的是地址 或 引用。所以上面整个示例是不太好的