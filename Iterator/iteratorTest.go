package Iterator

import "fmt"

func IteratorTest()  {
	ss:=[]string{"a","b","c"}
	ca:=&ConcreteAggregate{
		List:ss,
	}
	it:=ca.CreateIterator()
	for it.First(); !it.IsDone();it.Next()  {
		fmt.Printf("item:%v\n",it.CurrentItem())
	}

	pm:=&PayManager{
		List:make([]PayModel,0),
	}
	pm.CalcPay()
	it1:=pm.CreateIterator()
	for it1.First();!it1.IsDone();it1.Next(){
		item:=it1.CurrentItem().(PayModel)
		fmt.Printf("username:%s,pay:%.2f\n",item.UserName,item.Pay)
	}

	it2:=NewPayFilterIterator(pm.List)
	fmt.Println("只查看工资5000以上的员工的薪资情况")
	for it2.First(); !it2.IsDone(); it2.Next() {
		item:=it2.CurrentItem().(PayModel)
		fmt.Printf("username:%s,pay:%.2f\n",item.UserName,item.Pay)
	}

}
