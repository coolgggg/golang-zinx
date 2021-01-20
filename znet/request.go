package znet

import "github.com/aceld/golang-zinx/ziface"

type Request struct {
	//已经和客户端建立的链接
	conn ziface.IConnection

	//客户端请求的数据
	msg ziface.IMessage
}

//得到的链接
func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

//得到的请求数据
func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

func (r *Request) GetMsgId() uint32 {
	return r.msg.GetMsgId()
}
