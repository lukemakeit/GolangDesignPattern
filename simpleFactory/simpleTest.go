package simpleFactory

func SimpleTest() {
	f1:=Factory{}
	api1:=f1.CreateApi(1)
	api1.Operation("nice")

	api2:=f1.CreateApi(2)
	api2.Operation("good")
}
