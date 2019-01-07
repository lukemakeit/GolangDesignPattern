package TemplateMethod

import "fmt"

//对用户来说 他只需要使用这个接口即可
type SystemLogin interface {
	Login(lm *LoginModel) bool
}

//登录时,用户输入的信息
type LoginModel struct {
	Email string
	Password string
}
//需要子类实现的函数接口
type NeedImplApi interface {
	FindLoginUser(email string) *LoginModel
	EncryptPwd(pwd string) string
}
//登录模板,也就是父类
type LoginTemplate struct {
	Impl NeedImplApi
}
//EncryptPwd默认实现,后面普通用户登录 子类直接继承默认实现;员工登录,员工函数需要重新实现该函数
func (lt *LoginTemplate)EncryptPwd(pwd string) string {
	return pwd
}
//登录函数
func (lt *LoginTemplate)Login(lm *LoginModel) bool {
	//调用子类的FindLoginUser,普通用户 员工保存的用户数据不在一个表中,所以两者都需要实现FindLoginUser函数
	dblm:=lt.Impl.FindLoginUser(lm.Email)
	if dblm==nil{
		//数据库中没这个用户
		fmt.Printf("数据库中没有用户:%s\n",lm.Email)
		return false
	}
	encryptPwd:=lt.Impl.EncryptPwd(lm.Password)
	lm.Password=encryptPwd
	return lt.Match(lm,dblm)
}
//邮件 密码尝试匹配
func (lt *LoginTemplate)Match(lm *LoginModel,dblm *LoginModel) bool  {
	fmt.Printf("lm.email=%s lm.password=%s dblm.email=%s dblm.password=%s\n",lm.Email,lm.Password,dblm.Email,dblm.Password)
	if lm.Password == dblm.Password && lm.Email==dblm.Email {
		return true
	}
	return false
}

//普通用户登录
type NormalLogin struct {
	*LoginTemplate //用组合方式继承LoginTemplate, 在NormalLogin中我们只实现FindLoginUser 而EncryptPwd用LoginTemplate的实现
}
func (nl *NormalLogin)FindLoginUser(email string) *LoginModel {
	//我们模拟在数据库中执行select * from tb_users where email=?
	fmt.Printf("SQL:select * from tb_users where email=%q\n",email)
	usersMap:=make(map[string]*LoginModel,3)
	u1:=&LoginModel{
		"u1@qq.com",
		"1234",
	}
	u2:=&LoginModel{
		"u2@qq.com",
		"3456",
	}
	u3:=&LoginModel{
		"u3@qq.com",
		"5678",
	}
	usersMap["u1@qq.com"]=u1
	usersMap["u2@qq.com"]=u2
	usersMap["u3@qq.com"]=u3

	ret,ok:=usersMap[email]
	if ok==false {
		return nil
	}
	return ret
}
func NewNormalLogin() SystemLogin  {
	nl:=&NormalLogin{}
	lt:=&LoginTemplate{
		nl,
	}
	nl.LoginTemplate=lt
	return lt
}

//员工登录
type WorkerLogin struct {
	*LoginTemplate //用组合方式继承LoginTemplate, 在WorkerLogin中需要实现FindLoginUser EncryptPwd
}
func (wl *WorkerLogin)FindLoginUser(email string) *LoginModel {
	//我们模拟在数据库中执行select * from tb_workers where email=?
	fmt.Printf("SQL:select * from tb_workers where email=%q\n",email)
	workersMap:=make(map[string]*LoginModel,3)
	w1:=&LoginModel{
		"w1@qq.com",
		"1234",
	}
	w2:=&LoginModel{
		"w2@qq.com",
		"3456",
	}
	w3:=&LoginModel{
		"w3@qq.com",
		"5678",
	}
	workersMap["w1@qq.com"]=w1
	workersMap["w2@qq.com"]=w2
	workersMap["w3@qq.com"]=w3

	ret,ok:=workersMap[email]
	if ok==false {
		return nil
	}
	return ret
}
//工作人员的密码是经过加密的
func (wl *WorkerLogin)EncryptPwd(pwd string) string {
	fmt.Println("员工密码，使用MD5进行加密")
	return pwd
}
func NewWorkerLogin() SystemLogin  {
	wl:=&WorkerLogin{}
	lt:=&LoginTemplate{
		wl,
	}
	wl.LoginTemplate=lt
	return lt
}
