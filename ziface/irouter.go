package ziface

/**
路由抽象接口
路由里的数据都是IRequest
*/
type IRouter interface {
	//在处理conn业务前的狗子方法 hook
	PreHandle(request IRequest)

	//在处理conn业务的主方法
	Handle(request IRequest)

	//在处理conn业务之后的狗子方法 hook
	PostHandle(request IRequest)
}
