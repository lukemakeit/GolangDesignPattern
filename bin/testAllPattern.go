package main

import (
	"github.com/lukexwang/GolangDesignPattern/Builder"
	"github.com/lukexwang/GolangDesignPattern/Factory"
	"github.com/lukexwang/GolangDesignPattern/Prototye"
	"github.com/lukexwang/GolangDesignPattern/simpleFactory"
)

func main() {
	simpleFactory.SimpleTest()
	Factory.FactoryTest()
	Builder.OutoutBuilderTest()
	Prototype.PrototypeTest()
}
