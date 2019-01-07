package Bridge

import "fmt"

//抽象消息对象
type AbstractMessage interface {
	SendMessage(message string,toUser string)
}
//普通消息的实现
type CommonMessage struct {
	sender MessageSendApi
}
func (cm *CommonMessage)SendMessage(message string,toUser string)  {
	cm.sender.Send(message,toUser)
}
func NewCommonMsg(ms MessageSendApi) AbstractMessage  {
	return &CommonMessage{
		sender:ms,
	}
}
//加急消息的实现
type UrgencyMessage struct {
	sender MessageSendApi
}
func (um *UrgencyMessage)SendMessage(message string,toUser string)  {
	message="加急: "+message
	um.sender.Send(message,toUser)
}
func NewUrgencyMsg(ms MessageSendApi) AbstractMessage  {
	return &UrgencyMessage{
		sender:ms,
	}
}

//发送消息统一接口
type MessageSendApi interface {
	//发送消息
	Send(message string,toUser string)
}

//以站内短消息方式发送
type MessageSMS struct {}
func (ms *MessageSMS)Send(message string,toUser string)  {
	fmt.Println("使用站内短消息方式,发送消息:'"+message+"' 给"+toUser)
}
func NewSMS() MessageSendApi {
	return &MessageSMS{}
}
//以E-mail方式发送消息
type MessageEmail struct {}
func (ms *MessageEmail)Send(message string,toUser string)  {
	fmt.Println("使用E-mail方式,发送消息:'"+message+"' 给"+toUser)
}
func NewEmail() MessageSendApi {
	return &MessageEmail{}
}