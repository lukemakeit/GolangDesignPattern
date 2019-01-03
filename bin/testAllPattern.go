package main

import (
	"github.com/lukexwang/GolangDesignPattern/Composite2"
	"github.com/lukexwang/GolangDesignPattern/Factory"
	"github.com/lukexwang/GolangDesignPattern/Flyweight"
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
	//Composite1.CompositeTest()
	Composite2.CompositeTest()
	Flyweight.FlyWeightTest()
}
