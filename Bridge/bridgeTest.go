package Bridge

func BridgeTest()  {
	//站内短消息发送
	sms:=NewSMS()
	//普通消息对象
	cmsg:=NewCommonMsg(sms)
	cmsg.SendMessage("请喝一杯茶","luke")

	//紧急消息对象
	umsg:=NewUrgencyMsg(sms)
	umsg.SendMessage("你有一个告警","vito")

	//切换为邮件发送
	email:=NewEmail()
	//紧急消息对象
	umsg1:=NewUrgencyMsg(email)
	umsg1.SendMessage("你的机器有异常登录","luke")
}
