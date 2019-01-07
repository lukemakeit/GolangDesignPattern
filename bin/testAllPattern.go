package main

import (
	"github.com/lukexwang/GolangDesignPattern/Bridge"
	"github.com/lukexwang/GolangDesignPattern/Composite2"
	"github.com/lukexwang/GolangDesignPattern/Factory"
	"github.com/lukexwang/GolangDesignPattern/Observer"
	"github.com/lukexwang/GolangDesignPattern/Proxy"
	"github.com/lukexwang/GolangDesignPattern/Strategy"
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
	//Flyweight.FlyWeightTest()
	//Decorator.DecoratorTest()
	Bridge.BridgeTest()
	//Mediator.MediatorTest()
	Observer.ObserverTest()
	//Command.CalculaterTest()
	//Iterator.IteratorTest()
	//TemplateMethod.TemplateTest()
	Strategy.StrategyTest()
}
