package State

func StateTest()  {
	vm:=NewVoteManager()
	for i:=0;i<11; i++{
		vm.Vote("张三","11号小朋友很可爱")
	}
}