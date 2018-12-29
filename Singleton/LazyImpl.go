package Singleton

import "sync"

type example struct {
	Host string `json:"host"`
	Port int `json:"port"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

var instance *example
var mux sync.Mutex

// 双重检查加锁机制
//但是这种实现依然可能有问题 https://www.zhihu.com/question/35268028
func GetInstance() *example {
	if instance == nil{
		mux.Lock()
		defer mux.Unlock()
		if instance == nil{
			instance=new(example)
			readConfig(instance)
		}
	}
	return instance
}
func readConfig(ins *example)  {
	ins.UserName="luke"
	ins.Password="123"
	ins.Host="1.1.1.1"
	ins.Port=9527
}