package TemplateMethod

import "fmt"

func TemplateTest()  {
	hd:=NewHttpDownload("www.baidu.com","F盘://唐诗三百首")
	hd.Download()

	fd:=NewFtpDownload("www.google.com","E盘://下载")
	fd.Download()

	u1:=&LoginModel{
		"u1@qq.com",
		"1234",
	}
	normalUsrLogin:=NewNormalLogin()
	sucessOrNot:=normalUsrLogin.Login(u1)
	fmt.Printf("%s 登录成功了么？%v\n",u1.Email,sucessOrNot)

	w1:=&LoginModel{
		"w1@qq.com",
		"1234567",
	}
	workerLogin:=NewWorkerLogin()
	sucessOrNot=workerLogin.Login(w1)
	fmt.Printf("%s 登录成功了么？%v\n",w1.Email,sucessOrNot)
}
