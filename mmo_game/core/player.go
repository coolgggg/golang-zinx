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

//tp1 世界聊天， tp2坐标， tp3 动作， tp4 移动之后的坐标信息更新

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

//同步玩家上线的位置信息
func (p *Player) SyncSurrounding() {
	//获取当前玩家周围的玩家有哪些
	pids := WorldMgrObj.AoiMgr.GetPidsByPos(p.X, p.Z)
	players := make([]*Player, 0, len(pids))
	for _, pid := range pids {
		players = append(players, WorldMgrObj.GetPlayerByPid(int32(pid)))
	}

	//将当前玩家的位置信息通过 msg id 200 发送给周围玩家（让其他玩家看到自己）
	proto_msg := &pb.BroadCast{
		Pid: p.Pid,
		Tp:  2, //tp2 代表广播坐标
		Data: &pb.BroadCast_P{
			P: &pb.Position{
				X: p.X,
				Y: p.Y,
				Z: p.Z,
				V: p.V,
			},
		},
	}
	for _, player := range players {
		player.SendMsg(200, proto_msg)
	}

	//将周围的全部玩家位置信息 msg id 202 发送给当前玩家的客户端（让自己看到其他玩家）
	//制作一个pb.Player slice
	players_proto_msg := make([]*pb.Player, 0, len(players))
	for _, player := range players {
		//制作一个 message Player
		p := &pb.Player{
			Pid: player.Pid,
			P: &pb.Position{
				X: player.X,
				Y: player.Y,
				Z: player.Z,
				V: player.V,
			},
		}
		players_proto_msg = append(players_proto_msg, p)
	}

	//封装 SyncPlayers protobuf 数据
	SyncPlayers_proto_msg := &pb.SyncPlayers{
		Ps: players_proto_msg[:],
	}
	//发送
	p.SendMsg(202, SyncPlayers_proto_msg)

}

//广播当前玩家的位置移动信息
func (p *Player) UpdatePos(x, y, z, v float32) {
	//更新当前玩家player对象的坐标
	p.X = x
	p.Y = y
	p.Z = z
	p.V = v

	//组建广播proto协议 msg id 200，tp=4
	proto_msg := &pb.BroadCast{
		Pid: p.Pid,
		Tp:  4,
		Data: &pb.BroadCast_P{
			P: &pb.Position{
				X: p.X,
				Y: p.Y,
				Z: p.Z,
				V: p.V,
			},
		},
	}

	//获取当前玩家的周边玩家AOI九宫格之内的玩家
	players := p.GetSurroundingPlayers()

	//一次给每个玩家对应的客户端发送当前玩家的位置更新信息
	for _, player := range players {
		player.SendMsg(200, proto_msg)
	}
}

//获取当前玩家的周边玩家AOI九宫格之内的玩家
func (p *Player) GetSurroundingPlayers() []*Player {
	//得到周边玩家pid
	pids := WorldMgrObj.AoiMgr.GetPidsByPos(p.X, p.Z)

	//将所有的pid对应的player放到players切片中
	players := make([]*Player, 0, len(pids))

	for _, pid := range pids {
		players = append(players, WorldMgrObj.GetPlayerByPid(int32(pid)))
	}

	return players
}
