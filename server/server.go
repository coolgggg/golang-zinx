package main

import (
	"fmt"
	"github.com/aceld/golang-zinx/ziface"
	"github.com/aceld/golang-zinx/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

//func (this *PingRouter) PreHandle(request ziface.IRequest) {
//	fmt.Println("call router pre handle")
//	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping router\n"))
//	if err != nil {
//		fmt.Println("call back before ping error")
//	}
//}

func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("call ping router handle")
	//先读取客户端的数据，再会写ping.. ping.. ping
	fmt.Println("recv from client:msgI = ", request.GetMsgId(),
		", data=", string(request.GetData()))

	err := request.GetConnection().SendMsg(200, []byte("ping..ping..ping"))
	if err != nil {
		fmt.Println("resend msg err", err)
	}

	//_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping.. ping .. ping\n"))
	//if err != nil {
	//	fmt.Println("call back ping error")
	//}
}

//func (this *PingRouter) PostHandle(request ziface.IRequest) {
//	fmt.Println("call router after handle")
//	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping router\n"))
//	if err != nil {
//		fmt.Println("call back after ping error")
//	}
//}

type HelloZinxRouter struct {
	znet.BaseRouter
}

func (this *HelloZinxRouter) Handle(request ziface.IRequest) {
	fmt.Println("call hello router handle")
	//先读取客户端的数据，再会写ping.. ping.. ping
	fmt.Println("recv from client:msgI = ", request.GetMsgId(),
		", data=", string(request.GetData()))

	err := request.GetConnection().SendMsg(201, []byte("hello.. hello.. hello"))
	if err != nil {
		fmt.Println("resend msg err", err)
	}

	//_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping.. ping .. ping\n"))
	//if err != nil {
	//	fmt.Println("call back ping error")
	//}
}

//创建连接hook钩子函数
func DoConnectionBegin(conn ziface.IConnection) {
	fmt.Println("=====> DoConnectionBegin is called..")
	if err := conn.SendMsg(202, []byte("DoConnectionBegin BEGIN")); err != nil {
		fmt.Println(err)
	}

	//给当前链接设置一些属性
	fmt.Println("set conn name home..")
	conn.SetProperty("name", "老司机-在线开车")
	conn.SetProperty("home", "amoy")
}

//断开连接hook钩子函数
func DoConnectionLost(conn ziface.IConnection) {
	fmt.Println("=====> DoConnectionLost is called..")
	fmt.Println("conn ID=", conn.GetConnID(), "is lost")

	//获取链接属性
	if name, err := conn.GetProperty("name"); err == nil {
		fmt.Println("get property name=", name)
	}
	if home, err := conn.GetProperty("home"); err == nil {
		fmt.Println("get property home=", home)
	}
}

func main() {
	//创建一个server句柄，使用zinx的api
	s := znet.NewServer("[zinx]")

	//注册连接hook钩子函数
	s.SetOnConnStart(DoConnectionBegin)
	s.SetOnConnStop(DoConnectionLost)

	//给当前zinx框架添加自定义的router
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloZinxRouter{})

	//启动server
	s.Server()
}
