package Singleton

import "fmt"

func SinletonTest()  {
	s:=GetInstance()
	fmt.Printf("username=%s\n",s.UserName)

	s2:=GetInstance2()
	fmt.Printf("password=%s\n",s2.Password)

	s3:=GetInstance3()
	fmt.Printf("host=%s port=%d\n",s3.Host,s3.Port)
}
