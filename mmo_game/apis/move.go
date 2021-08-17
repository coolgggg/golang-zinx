package apis

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"golang-zinx/mmo_game/core"
	"golang-zinx/mmo_game/pb"
	"golang-zinx/ziface"
	"golang-zinx/znet"
)

//玩家移动
type MoveApi struct {
	znet.BaseRouter
}

func (m *MoveApi) Handle(request ziface.IRequest) {
	//解析客户端传递进来的proto协议
	proto_msg := &pb.Position{}
	err := proto.Unmarshal(request.GetData(), proto_msg)
	if err != nil {
		fmt.Println("move: position unmarshal err ", err)
	}
	//得到当前发送位置的是哪个玩家
	pid, err := request.GetConnection().GetProperty("pid")
	if err != nil {
		fmt.Println("get property err: ", err)
		return
	}

	fmt.Println("player pid = %d, move (%f, %f, %f, %f)\n", pid, proto_msg.X, proto_msg.Y, proto_msg.Z, proto_msg.V)

	//给其他玩家进行当前玩家的位置信息广播
	player := core.WorldMgrObj.GetPlayerByPid(pid.(int32))

	//广播并更新当前玩家的坐标
	player.UpdatePos(proto_msg.X, proto_msg.Y, proto_msg.Z, proto_msg.V)

}
