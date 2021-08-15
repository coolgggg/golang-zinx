package main

import (
	"fmt"
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

	fmt.Println("====>player pid =", player.Pid, " is arrived <====")
}