package simpleFactory

import (
	"fmt"
)

type API interface {
	Operation(s string)
}

type ImplA struct {}

func (i ImplA)Operation(s string)  {
	fmt.Println("ImplA s=="+s)
}
type ImplB struct {}

func (i ImplB)Operation(s string)  {
	fmt.Println("ImplB s=="+s)
}

type Factory struct {}

func (f Factory)CreateApi(condition int) API {
	var api API
	if condition == 1 {
		api = ImplA{}
	}else if condition ==2 {
		api = ImplB{}
	}
	return api
}
