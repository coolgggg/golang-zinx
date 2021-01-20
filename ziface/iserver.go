package ziface

type IServer interface {
	Start()

	Stop()

	Server()

	//路由功能，给当前的服务器注册一个路由分发，供客户端的链接处理使用
	AddRouter(msgId uint32, router IRouter)

	//获取当前server的连接管理器
	GetConnMgr() IConnManager

	//注册 OnConnStart 钩子函数的方法
	SetOnConnStart(func(connection IConnection))
	//注册 OnConnStop 钩子函数的方法
	SetOnConnStop(func(connection IConnection))
	//调用 OnConnStart 钩子函数的方法
	CallOnConnStart(connection IConnection)
	//调用 OnConnStop 钩子函数的方法
	CallOnConnStop(connection IConnection)
}
