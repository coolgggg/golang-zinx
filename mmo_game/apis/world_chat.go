package apis

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"golang-zinx/mmo_game/core"
	"golang-zinx/mmo_game/pb"
	"golang-zinx/ziface"
	"golang-zinx/znet"
)

//世界聊天 路由业务
type WorldChatApi struct {
	znet.BaseRouter
}

func (wc *WorldChatApi) Handle(request ziface.IRequest) {
	//解析客户端传递进来的proto协议
	proto_msg := &pb.Talk{}
	err := proto.Unmarshal(request.GetData(), proto_msg)
	if err != nil {
		fmt.Println("talk unmarshal err ", err)
	}

	//当前的聊天数据是哪个玩家发送的
	pid, err := request.GetConnection().GetProperty("pid")
	if err != nil {
		fmt.Println("get pid err ", err)
	}

	//根据pid得到player对象
	player := core.WorldMgrObj.GetPlayerByPid(pid.(int32))

	//将这个消息广播给全部在线的玩家
	player.Talk(proto_msg.Content)
}
