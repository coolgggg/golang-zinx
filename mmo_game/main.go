package main

import (
	"fmt"
	"golang-zinx/mmo_game/apis"
	"golang-zinx/mmo_game/core"
	"golang-zinx/ziface"
	"golang-zinx/znet"
)

func main() {
	//创建zinx server 句柄
	s := znet.NewServer("MMO Game Zinx Cool")

	//连接创建和销毁的Hook钩子
	s.SetOnConnStart(OnConnectionAdd)

	//注册一些路由业务
	s.AddRouter(2, &apis.WorldChatApi{})
	s.AddRouter(3, &apis.MoveApi{})

	s.Server()
}

//创建连接hook钩子函数
func OnConnectionAdd(conn ziface.IConnection) {
	//创建一个player对象
	player := core.NewPlayer(conn)

	//给客户端发送msg id：1 的消息，同步当前player的id给客户端
	player.SyncPid()

	//给客户端发送msg id：200 的消息，同步当前位置给客户端
	player.BroadCastStartPosition()

	//将当前新上线的玩家添加到world mgr
	core.WorldMgrObj.AddPlayer(player)

	//将该连接绑定一个pid
	conn.SetProperty("pid", player.Pid)

	//同步周边玩家，告知他们当前玩家已经上线，广播当前玩家的消信息
	player.SyncSurrounding()

	fmt.Println("====>player pid =", player.Pid, " is arrived <====")
}

//给当前链接断开之前触发的hook钩子函数
func OnConnectionLost(conn ziface.IConnection) {
	//通过属性拿pid、信息
	pid, _ := conn.GetProperty("pid")
	player := core.WorldMgrObj.GetPlayerByPid(pid.(int32))

	//在连接断开之前处理玩家下线业务
	player.Offline()

	fmt.Println("===> player pid =", pid, "offline ... ")
}
