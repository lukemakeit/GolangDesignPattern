package Proxy

import (
	"fmt"
)

//用户数据对象接口
type UserModelApi interface {
	GetUserId() string
	SetUserId(userId string)
	GetName() string
	SetName(name string)
	GetDepId() string
	SetDepId(depId string)
	GetSex() string
	SetSex(sex string)
}

//用户数据对象
type UserModel struct {
	userID string `json:"userId"`
	name string `json:"name"`
	depID string `json:"depId"`
	sex string `json:"sex"`
}
func (um *UserModel)SetUserId(userId string)  {
	um.userID=userId
}
func (um *UserModel)GetUserId() string  {
	return um.userID
}
func (um *UserModel)SetName(name string)  {
	um.name=name
}
func (um *UserModel)GetName() string {
	return um.name
}
func (um *UserModel)SetDepId(depId string)  {
	um.depID=depId
}
func (um *UserModel)GetDepId() string {
	return um.depID
}
func (um *UserModel)SetSex(sex string)  {
	um.sex=sex
}
func (um *UserModel)GetSex() string {
	return um.sex
}

type UserProxy struct {
	RealUser UserModel
	Loaded bool //是否已经重新装载过数据
}
func (up *UserProxy)SetUserId(userId string)  {
	up.RealUser.SetUserId(userId)
}
func (up *UserProxy)GetUserId() string  {
	return  up.RealUser.GetUserId()
}
func (up *UserProxy)SetName(name string)  {
	up.RealUser.SetName(name)
}
func (up *UserProxy)GetName() string {
	return up.RealUser.GetName()
}
func (up *UserProxy)SetDepId(depId string)  {
	up.RealUser.SetDepId(depId)
}
func (up *UserProxy)GetDepId() string {
	//延迟加载,等真正需要这个用户的DepId信息时再加载
	if !up.Loaded{
		up.reload()
		up.Loaded=true
	}
	return up.RealUser.GetDepId()
}
func (up *UserProxy)SetSex(sex string)  {
	up.RealUser.SetSex(sex)
}
func (up *UserProxy)GetSex() string {
	//延迟加载,等真正需要这个用户的sex信息时再加载
	if !up.Loaded{
		up.reload()
		up.Loaded=true
	}
	return up.RealUser.GetSex()
}
func (up *UserProxy)reload()  {
	fmt.Printf("重新查询数据库获取完整用户数据,userId==%s\n",up.RealUser.GetUserId())
	fmt.Printf("Connect DB and Run SQL:select * from tbl_user where userId='%s';\n",up.RealUser.GetUserId())
	//我们这里就随便模拟一下即可
	up.RealUser.SetDepId("9999")
	up.RealUser.SetSex("Male")
}

//UserManaager
//从数据库查询值时,不需要获取全部字段,只需要查询得到用户编号、姓名即可
//在数据库中获取的值转变成对象时,创建的对象不再是UserModel,而是代理对象
type UserManager struct {}
func (um *UserManager)GetUserByDepId(depId string) ([]UserModelApi) {
	users:=[]UserModelApi{}
	//这里就随便模拟下查询数据库,得到三条数据
	fmt.Printf("Ok,Now query data about department:%s\n",depId)
	fmt.Printf("SQL:select u.userId,u.name from tbl_user u, tbl_dep d where u.depId=d.depId and d.depId like '%s%s'\n",depId,"%")
	u1:=UserProxy{
		RealUser:UserModel{},
	}
	u1.SetUserId("0010101")
	u1.SetName("张三")

	u2:=UserProxy{
		RealUser:UserModel{},
	}
	u2.SetUserId("0010102")
	u2.SetName("李四")

	u3:=UserProxy{
		RealUser:UserModel{},
	}
	u3.SetUserId("0010103")
	u3.SetName("王五")

	users=append(users,&u1)
	users=append(users,&u2)
	users=append(users,&u3)
	return  users
}
