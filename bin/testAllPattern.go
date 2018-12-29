package main

import (
	"github.com/lukexwang/GolangDesignPattern/Factory"
	"github.com/lukexwang/GolangDesignPattern/Proxy"
	"github.com/lukexwang/GolangDesignPattern/simpleFactory"
)

func main() {
	simpleFactory.SimpleTest()
	Factory.FactoryTest()
	//Builder.OutoutBuilderTest()
	//Prototype.PrototypeTest()
	//Singleton.SinletonTest()
	//Facade.FacadeTest()
	//Adapter.AdapterTest()
	Proxy.ProxyTest()
}
