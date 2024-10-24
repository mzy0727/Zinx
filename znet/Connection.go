package znet

import (
	"fmt"
	"net"
	"zinx/ziface"
)

type Connection struct {
	// 当前连接的socket tcp套接字
	Conn *net.TCPConn
	// 连接的id
	ConnID uint32
	// 当前的连接状态
	isClosed bool
	// 当前链接所绑定的处理业务方法的api
	handleAPI ziface.HandleFunc
	// 告知当前连接已经退出/停止 channel
	ExitChan chan bool
}

// 初始化连接模块的方法
func NewConnection(conn *net.TCPConn, connID uint32, callback_api ziface.HandleFunc) *Connection {
	c := &Connection{
		Conn:      conn,
		ConnID:    connID,
		isClosed:  false,
		handleAPI: callback_api,
		ExitChan:  make(chan bool, 1),
	}

	return c
}

// 连接的读业务方法
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running...")
	defer fmt.Println("ConnID =", c.ConnID, "Reader is exit,remote addr is", c.Conn.RemoteAddr())
	defer c.Stop()

	for {
		// 读取客户端的数据到buf中
		buf := make([]byte, 1024)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("Reader Goroutine is exit,err:", err)
			continue
		}

		// 调用当前连接所绑定的handleapi
		if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
			fmt.Println("ConnID:", c.ConnID, " Handle is error,err:", err)
		}
	}
}

func (c *Connection) Start() {
	fmt.Println("Conn Start()... ConnID=", c.ConnID)
	// 启动从当前链接的读数据的业务
	go c.StartReader()
	// TODO 启动从当前链接的写数据的业务
}
func (c *Connection) Stop() {
	fmt.Println("Conn stop()... ConnID=", c.ConnID)

	if c.isClosed == true {
		return
	}
	c.isClosed = true
	// 关闭socket连接
	c.Conn.Close()
	// 回收资源
	close(c.ExitChan)
}

func (c *Connection) GetTcpConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) Send(data []byte) error {
	return nil
}
