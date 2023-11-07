package im

import (
	"IM-Service/src/configs/conf"
	"IM-Service/src/configs/log"
	"IM-Service/src/service"
	"im-sdk/client"
	"im-sdk/handler"
	"im-sdk/model"
)

type LogicProcess struct{}

var process = &LogicProcess{}

func GetLogicProcess() *LogicProcess {
	return process
}
func StartIM() {
	//启动长连接
	conf.Conf.Connected = true
	conf.Conf.Client = client.New(conf.Base.WsHost)
	e := conf.Conf.Client.Startup(GetLogicProcess())
	if e != nil {
		_ = log.WithError(e, "启动长连接失败")
		conf.Conf.Connected = false
	}
	go func() {
		mgr := handler.GetClientHandler().GetMessageManager()
		//开启心跳
		mgr.StartupHeartbeat()
		//开启Qos
		mgr.StartupQos()
	}()
}

// SendOk qos中的消息发送成功 服务器成功返回
func (_self *LogicProcess) SendOk(protocol *model.Protocol) {
	messageService := service.NewMessageService()
	messageService.UpdateReaded(protocol, 2)
	service.Listener.OnSendReceive([]byte("123"))
}

// SendOkCallback 发送成功的回调
// 仅仅是发出去了 如果是Qos消息 此时还未收到服务器反馈
// SendOk代表发出Qos消息并接收到了服务器反馈
func (_self *LogicProcess) SendOkCallback(protocol *model.Protocol) {

}

// SendFailedCallback 发送失败的回调
func (_self *LogicProcess) SendFailedCallback(protocol *model.Protocol) {
	messageService := service.NewMessageService()
	messageService.UpdateReaded(protocol, -1)
	//api.Listener.OnSendReceive()
}

// LoginOk 登录成功的回调
func (_self *LogicProcess) LoginOk(protocol *model.Protocol) {
	log.Debugf("登录成功！:%v", protocol)
	//获取离线消息
	e := service.NewMessageService().GetOfflineMessage()
	if e != nil {
		log.Error(e)
	}
}

// LoginFail 登录失败的回调
func (_self *LogicProcess) LoginFail(protocol *model.Protocol) {
	log.Debugf("登录失败！:%v", protocol)
}

// ReceivedMessage 接收到消息
func (_self *LogicProcess) ReceivedMessage(protocol *model.Protocol) {
	log.Debugf("接收到服务器IM消息:%v", protocol)
	err := service.NewMessageService().Handler(protocol)
	if err != nil {
		log.Errorf("解析服务器IM消息失败:%v", err)
	}
	//避免循环引用  这里单独处理事件
	if service.Listener != nil {
		service.Listener.OnReceive([]byte("123"))
	}
}
func (_self *LogicProcess) Exception(msg string) {
	log.Errorf("exception:%v", msg)
	if msg == "unexpected EOF" {
		log.Debug("服务器断开连接,进行重连")
		go func() {
			err := conf.Conf.Client.Reconnect()
			if err != nil {
				log.Error(err)
			}
		}()
	}
}