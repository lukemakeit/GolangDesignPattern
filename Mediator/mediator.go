package Mediator

import (
	"fmt"
	"strings"
)

//同事类
//光驱类
type CDDriver struct {
	data string
	media *Mediator
}
func (cd *CDDriver)GetData() string  {
	return cd.data
}
func (cd *CDDriver)ReadCD()  {
	cd.data="设计模式,值得好好研究"
	//通知主板,光驱中的数据已经读好了
	cd.media.Changed(cd)
}
func NewCDDriver(media *Mediator) *CDDriver {
	return &CDDriver{
		media:media,
	}
}
//CPU类
type CPU struct {
	video string
	sound string
	media *Mediator
}
func (c *CPU)GetVideoData() string  {
	return c.video
}
func (c *CPU)GetSoundData() string  {
	return c.sound
}
func (c *CPU)ExecuteData(data string)  {
	ss:=strings.Split(data,",")
	c.video=ss[0]
	c.sound=ss[1]
	//通知主板,CPU工作已完成
	c.media.Changed(c)
}
func NewCPU(media *Mediator) *CPU {
	return &CPU{
		media:media,
	}
}
//显卡类
type VideoCard struct {
	media *Mediator
}
func (vc *VideoCard)ShowData(data string)  {
	fmt.Printf("您正在观看的是:%s\n",data)
}
func NewVideoCard(media *Mediator) *VideoCard {
	return &VideoCard{
		media:media,
	}
}
//声卡类
type SoundCard struct {
	media *Mediator
}
func (sc *SoundCard)ShowData(data string)  {
	fmt.Printf("你正在听到:%s\n",data)
}
func NewSoundCard(media *Mediator) *SoundCard {
	return &SoundCard{
		media:media,
	}
}

//中介者类
type Mediator struct {
	CD *CDDriver
	CPU *CPU
	Video *VideoCard
	Sound *SoundCard
}
func (m *Mediator)Changed(i interface{})  {
	switch inst:=i.(type) {
	case *CDDriver:
		m.OpeCDDriverReadData(inst)
	case *CPU:
		m.OpeCPU(inst)
	}
}
//光驱读取数据后与其他对象交互
func (m *Mediator)OpeCDDriverReadData(cd *CDDriver)  {
	data:=cd.GetData()
	m.CPU.ExecuteData(data)
}
//处理CPU处理完数据后与其他对象交互
func (m *Mediator)OpeCPU(c *CPU)  {
	m.Video.ShowData(c.GetVideoData())
	m.Sound.ShowData(c.GetSoundData())
}