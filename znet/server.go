package znet

import (
	"fmt"
	"net"
	"zinx/utils"
	"zinx/ziface"
)

// iServer 接口的实现，定义一个Server服务类
type Server struct {
	// 服务器的名称
	Name string
	// IP版本号
	IPVersion string
	// 服务器绑定的IP地址
	IP string
	// 服务器绑定的端口
	port int
	// 当前的Server添加一个router，server注册的连接对应的处理业务
	Router ziface.IRouter
}

// 开启网络服务
func (s *Server) Start() {
	fmt.Printf("[Start] server listenner at IP : %s, Port : %d\n", s.IP, s.port)
	// 开启一个go去做服务端的listen业务
	go func() {
		// 1. 获取一个TCP的addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.port))
		if err != nil {
			fmt.Printf("[Start] resolve tcp addr err : %s", err.Error())
			return
		}
		// 2. 监听服务器地址
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Printf("[Start] listen tcp err : %s", err.Error())
			return
		}
		fmt.Println("[Start] server start success\n")

		var cid uint32
		cid = 0

		// 3. 阻塞等待客户端连接，处理客户端业务（读写）
		for {
			// 如果有客户端连接，阻塞会返回
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Printf("[Start] accept err : %s", err.Error())
				continue
			}

			// 将处理新连接的业务方法和conn绑定 得到连接模块
			dealConn := NewConnection(conn, cid, s.Router)
			cid++

			// 启动当前的连接业务处理
			go dealConn.Start()
		}
	}()

}

// 关闭网络服务
func (s *Server) Stop() {
	fmt.Println("[STOP] Zinx server, name : ", s.Name)
	// TODO Server.Stop()  将其他需要清理的连接信息或者其他信息 也要一并停止或者清理
}

// 运行网络服务
func (s *Server) Serve() {
	// 启动server的服务功能
	s.Start()

	// TODO 做一些启动服务器之后的额外的业务

	// 阻塞状态
	select {}
}

// 添加路由方法
func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
	fmt.Println("[AddRouter] Zinx server name : ", s.Name)
}

// 初始化服务器
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      utils.GlobalObject.Name,
		IPVersion: "tcp4",
		IP:        utils.GlobalObject.Host,
		port:      utils.GlobalObject.TcpPort,
		Router:    nil,
	}
	return s
}
