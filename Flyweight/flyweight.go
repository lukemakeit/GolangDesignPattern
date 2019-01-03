package Flyweight

import (
	"strings"
	"sync"
)

//授权数据享元接口
type Flyweight interface {
	Match(securityEntity string,permit string) bool
}
//授权数据中安全实体 权限组成享元对象
type AuthorizationFlyweith struct {
	securityEntity string
	permit string
}
func (af *AuthorizationFlyweith)GetSecurityEntity() string  {
	return af.securityEntity
}
func (af *AuthorizationFlyweith)GetPermit() string  {
	return af.permit
}
//实现享元接口
func (af *AuthorizationFlyweith)Match(securityEntity string,permit string) bool {
	if af.securityEntity== securityEntity && af.permit==permit {
		return true
	}
	return false
}
//享元对象构造方法
//@param state状态数据,包含安全实体和权限数据,用逗号分隔
func NewAuthorizationFlyweith(state string) *AuthorizationFlyweith  {
	ss:=strings.Split(state,",")
	return &AuthorizationFlyweith{
		securityEntity:ss[0],
		permit:ss[1],
	}
}

//享元工厂,通常实现为单例模式
var flyweightFactory map[string]Flyweight
var once sync.Once

func GetFlyWFactory() map[string]Flyweight {
	once.Do(func() {
		flyweightFactory=make(map[string]Flyweight)
	})
	return flyweightFactory
}
func GetFlyWeightByKey(key string) Flyweight  {
	f,ok:=GetFlyWFactory()[key]
	if ok==false{
		f=NewAuthorizationFlyweith(key)
		GetFlyWFactory()[key]=f
	}
	return f
}

//安全管理,管理每个用户的权限,单例模式
var securityMgr map[string][]Flyweight
var once1 sync.Once

func GetSecurityMgr() map[string][]Flyweight  {
	once1.Do(func() {
		//在初始化时,没有主动去数据库中拉取所有玩家的权限
		securityMgr=make(map[string][]Flyweight)
	})
	return securityMgr
}
//模拟从DB中查询数据
func QueryByUser(user string)[]Flyweight  {
	userAutFlys:=make([]Flyweight,0)
	for _,v:=range testDB.GetUserAndPrivs(){
		ss:=strings.Split(v,",")
		if ss[0] == user {
			f:=GetFlyWeightByKey(ss[1]+","+ss[2])
			userAutFlys=append(userAutFlys,f)
		}
	}
	return userAutFlys
}
//用户login时,获取其权限信息放到securityMgr中
func Login(user string)  {
	userAutFlys:=QueryByUser(user)
	GetSecurityMgr()[user]=userAutFlys
}
//判断某用户对某个安全实体是否有某权限
func HasPermit(user string,securityEntity string,permit string) bool  {
	userAutFlys,ok:=GetSecurityMgr()[user]
	if ok==false {
		return false
	}
	for _,u:=range userAutFlys{
		if u.Match(securityEntity,permit) {
			return true
		}
	}
	return false
}
