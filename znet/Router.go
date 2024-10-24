package znet

import "zinx/ziface"

// 实现router时，先嵌入这个baseRouter基类，然后根据需要对这个基类的方法进行重写
type BaseRouter struct {
}

// 这里之所以baseRouter都为空
// 是因为有的router 不希望有PreHandle,PostHandle这两个业务
// 所以router全部继承baserouter，实现了接口全部方法
func (br *BaseRouter) PreHandle(request ziface.IRequest)  {}
func (br *BaseRouter) Handle(request ziface.IRequest)     {}
func (br *BaseRouter) PostHandle(request ziface.IRequest) {}
