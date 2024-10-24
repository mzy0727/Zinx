package ziface

import "net"

type IConnection interface {
	// 启动连接
	Start()
	// 停止连接
	Stop()
	// 获取当前连接的绑定socket coon
	GetTcpConnection() *net.TCPConn
	// 获取当前连接模块的连接id
	GetConnID() uint32
	// 获取远程客户端的tcp状态 ip port
	GetRemoteAddr() net.Addr
	// 发送数据，将数据发送给远程的客户端
	Send(data []byte) error
}

// 定义一个处理连接业务的方法
type HandleFunc func(*net.TCPConn, []byte, int) error
