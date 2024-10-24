package znet

import (
	"fmt"
	"net"
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
}

// 定义当前客户端连接的所绑定的handle api(目前写死，以后优化，应用用户自定义）
func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	// 回显业务
	fmt.Println("[call back] to client")
	if _, err := conn.Write(data); err != nil {
		fmt.Println("write back buf err", err)
		return err
	}
	return nil
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
			dealConn := NewConnection(conn, cid, CallBackToClient)
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

// 初始化服务器
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		port:      7777,
	}
	return s
}
