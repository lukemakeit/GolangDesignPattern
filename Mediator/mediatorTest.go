package Mediator

func MediatorTest()  {
	media:=&Mediator{}
	cpu:=NewCPU(media)
	cd:=NewCDDriver(media)
	video:=NewVideoCard(media)
	sound:=NewSoundCard(media)

	media.CPU=cpu
	media.CD=cd
	media.Video=video
	media.Sound=sound

	//光驱开始读取数据
	cd.ReadCD()
}
