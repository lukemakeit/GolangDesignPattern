package Observer

func ObserverTest()  {
	zs:=&Reader{
		Name:"张三",
	}
	ls:=&Reader{
		Name:"李四",
	}
	ww:=&Reader{
		Name:"王五",
	}
	np:=NewsPaper{
		Readers:make([]Observer,0),
		Content:"这是一份新报纸,感谢你的阅读",
	}
	np.Attach(zs)
	np.Attach(ls)
	np.Attach(ww)

	np.Notify()
}
