package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"zinx/ziface"
)

/*
	存储一切有关Zinx框架的全局参数，供其他模块使用
	一些参数是可以通过zinx.json由用户进行配置
*/

type GlobalObj struct {
	/*
		Server
	*/
	TcpServer ziface.IServer // 当前Zinx全局的Server对象
	Host      string         // 监听的ip
	TcpPort   int            // 端口号
	Name      string         // 服务器名称

	/*
		Zinx
	*/

	Version       string // 版本号
	MaxConn       int    // 最大连接数量
	MaxPacketSize uint32 // 数据包最大值
}

/*
定义一个全局的对外的Globalobj
*/
var GlobalObject *GlobalObj

// 从zinx.json去加载用于自定义的参数
func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		fmt.Println("read conf/zinx.json fail")
		panic(err)
	}
	// 将json文件参数解析到struct中
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		fmt.Println("parse conf/zinx.json fail")
		panic(err)
	}

}

/*
提供一个init方法，初始化当前的GlobalObject
*/
func init() {
	// 如果配置文件没有加载，默认的值
	GlobalObject = &GlobalObj{
		Name:          "Zinx",
		Version:       "v0.4",
		TcpPort:       7777,
		Host:          "0.0.0.0",
		MaxConn:       1000,
		MaxPacketSize: 4096,
	}

	// 应该尝试从conf/zinx.json去加载一些用户自定义的参数
	GlobalObject.Reload()
}
