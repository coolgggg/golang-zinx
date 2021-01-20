package znet

import "github.com/aceld/golang-zinx/ziface"

type BaseRouter struct {
}

//这里之所以base router的方法都为空
//是因为有的router可选实现需要的方法

//在处理conn业务前的狗子方法 hook
func (br *BaseRouter) PreHandle(request ziface.IRequest) {}

//在处理conn业务的主方法
func (br *BaseRouter) Handle(request ziface.IRequest) {}

//在处理conn业务之后的狗子方法 hook
func (br *BaseRouter) PostHandle(request ziface.IRequest) {}
