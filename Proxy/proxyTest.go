package Proxy

import "fmt"

func ProxyTest()  {
	um:=UserManager{}
	users:=um.GetUserByDepId("00101")
	fmt.Printf("userId:%s userName:%s\n",users[0].GetUserId(),users[0].GetName())
	fmt.Printf("userId:%s userName:%s\n",users[1].GetUserId(),users[1].GetName())
	fmt.Printf("userId:%s userName:%s depId:%s Sex:%s\n",users[2].GetUserId(),users[2].GetName(),users[2].GetDepId(),users[2].GetSex())
}
