package Singleton

type example2 struct {
	Host string `json:"host"`
	Port int `json:"port"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

var instance2 *example2

//包初始化时就执行,实例化单例
func init()  {
	instance2=new(example2)
	readConfig2(instance2)
}

func readConfig2(ins *example2)  {
	ins.UserName="luke"
	ins.Password="123"
	ins.Host="1.1.1.1"
	ins.Port=9527
}

func GetInstance2() *example2 {
	return  instance2
}