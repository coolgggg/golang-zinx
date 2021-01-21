package utils

import (
	"encoding/json"
	"golang-zinx/ziface"
	"io/ioutil"
)

//存储一切有关zinx框架的全局参数，供其他模块使用
//一些参数是可以通过zinx.json 由用户进行配置
type GlobalObj struct {

	//server
	TcpServer ziface.IServer //当前zinx全局的server对象
	Host      string
	TcpPort   int
	Name      string

	//zinx
	Version          string
	MaxConn          int    //当前服务器允许的最大连接数
	MaxPackageSize   uint32 //当前zinx框架数据包的最大值
	WorkerPoolSize   uint32 //当前业务工作worker池的goroutine数量
	MaxWorkerTaskLen uint32 //zinx框架允许用户最多开辟多少个worker（限定条件）
}

//定义一个全局的对外的global obj
var GlobalObject *GlobalObj

//从zinx.json 去加载用于自定义的参数
func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("../conf/zinx.json")
	if err != nil {
		//fmt.Println("read config json err")
		panic(err)
	}
	//将json文件数据解析到struct中
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

//提供一个init方法，初始化当前的GlobalObj
func init() {
	//如果配置文件没有价值，默认的值
	GlobalObject = &GlobalObj{
		Name:             "Zinx server app",
		Version:          "V0.6",
		TcpPort:          8999,
		Host:             "0.0.0.0",
		MaxConn:          1000,
		MaxPackageSize:   4096,
		WorkerPoolSize:   10,
		MaxWorkerTaskLen: 1024, //每个worker对应的消息队列的任务数量最大值
	}

	//尝试从 conf/zinx。json去加载一些用户自定义的参数
	GlobalObject.Reload()
}
