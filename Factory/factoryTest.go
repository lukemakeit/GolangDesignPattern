package Factory

func FactoryTest()  {
	bjS1:=BJBunShop{}
	bjPigB:=bjS1.Generate("pig")
	bjPigB.Eat()

	gdS1:=GDBunShop{}
	gdSunB:=gdS1.Generate("sun")
	gdSunB.Eat()
}
