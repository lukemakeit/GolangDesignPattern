package Singleton

import "sync"

type example3 struct {
	Host string `json:"host"`
	Port int `json:"port"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

var instance3 *example3
var once sync.Once

func GetInstance3() *example3  {
	once.Do(func() {
		instance3=new(example3)
		readConfig3(instance3)
	})
	return instance3
}
func readConfig3(ins *example3)  {
	ins.UserName="luke"
	ins.Password="123"
	ins.Host="1.1.1.1"
	ins.Port=9527
}
