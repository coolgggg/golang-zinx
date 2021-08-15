package core

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"golang-zinx/mmo_game/pb"
	"golang-zinx/ziface"
	"math/rand"
	"sync"
)

//玩家对象
type Player struct {
	Pid  int32              //玩家id
	Conn ziface.IConnection //当前玩家的连接（用于和客户端的连接）
	X    float32            //平面x坐标
	Y    float32            //高度
	Z    float32            //平面Y坐标
	V    float32            //旋转0-360角度
}

//player id 生成器
var PidGen int32 = 1  //用来生产玩家ID的计数器
var IdLock sync.Mutex //保护PidGen的

func NewPlayer(conn ziface.IConnection) *Player {
	//生成一个玩家ID
	IdLock.Lock()
	id := PidGen
	PidGen++
	IdLock.Unlock()

	//创建一个玩家对象
	p := &Player{
		Pid:  id,
		Conn: conn,
		X:    float32(160 + rand.Intn(10)), //随机在160坐标点，基于x轴若干偏移
		Y:    0,
		Z:    float32(140 + rand.Intn(20)), //随机在140坐标点，基于Y轴若干偏移
		V:    0,
	}
	return p
}

//提供一个发送给客户端消息的方法
//主要是将pb的protobuf序列化之后，再调用zinx的SendMsg方法
func (p *Player) SendMsg(msgId uint32, data proto.Message) {
	//将proto message结构体序列化，装换成二进制
	msg, err := proto.Marshal(data)
	if err != nil {
		fmt.Println("marshal msg err: ", err)
	}

	//将二进制文件，通过zinx框架的SendMsg将数据发送给客户端
	if p.Conn == nil {
		fmt.Println("connection in players is nil")
	}

	if err := p.Conn.SendMsg(msgId, msg); err != nil {
		fmt.Println("Player SendMsg error! ")
		return
	}
	return
}

//告知客户端玩家PId，同步一截生成的玩家id给客户端
func (p *Player) SyncPid() {
	//组件msg id:1 的protobuf数据
	data := &pb.SyncPid{
		Pid: p.Pid,
	}

	//将消息发送给客户端
	p.SendMsg(1, data)
}

//广播玩家的出生地点
func (p *Player) BroadCastStartPosition() {
	//组件msg id:200 的protobuf数据
	msg := &pb.BroadCast{
		Pid: p.Pid,
		Tp:  2, //Tp2代表广播的位置坐标
		Data: &pb.BroadCast_P{
			P: &pb.Position{
				X: p.X,
				Y: p.Y,
				Z: p.Z,
				V: p.V,
			},
		},
	}

	p.SendMsg(200, msg)
}

//玩家广播世界聊天消息
func (p *Player) Talk(content string) {
	//组建msg id 200 的数据
	proto_msg := &pb.BroadCast{
		Pid: p.Pid,
		Tp:  1, //type=1代表聊天广播
		Data: &pb.BroadCast_Content{
			Content: content,
		},
	}

	//得到当前在线的玩家
	players := WorldMgrObj.GetAllPlays()

	//向所有玩家（包括自己）发送消息
	for _, player := range players {
		//给每个客户端发消息
		player.SendMsg(200, proto_msg)
	}

}
