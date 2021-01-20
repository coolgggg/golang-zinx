package ziface

/*
 IRequest 接口
实际上是吧客户端请求的链接信息 和 请求的数据  包装到了一个request中
*/
type IRequest interface {
	//得到的链接
	GetConnection() IConnection

	//得到的请求数据
	GetData() []byte

	//得到请求的id
	GetMsgId() uint32
}
