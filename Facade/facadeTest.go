package Facade

import "fmt"

func FacadeTest()  {
	ft:=NewSystemMoudleAPI()
	fmt.Println(ft.Test())
}
