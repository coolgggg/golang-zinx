package znet

import (
	"fmt"
	"github.com/aceld/golang-zinx/utils"
	"github.com/aceld/golang-zinx/ziface"
	"net"
)

//iServer的接口实现
type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
	//当前的server添加一个router, server 注册的链接对应的业务
	//Router ziface.IRouter
	//当前server的消息管理模块，用来绑定msgId和读研业务api关系
	MsgHandler ziface.IMsgHandle
	//该server的连接管理器
	ConnMgr ziface.IConnManager

	//hook
	OnConnStart func(conn ziface.IConnection)
	OnConnStop  func(conn ziface.IConnection)
}

/*
//定义当前客户端链接锁绑定的handle api，目前这个handle是写死的，以后优化应用去自定义
func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	//回显的业务
	fmt.Println("[conn handle] callback to client ... ")
	_, err := conn.Write(data[:cnt])
	if err != nil {
		fmt.Println("write back buf err", err)
		return errors.New("callback to client error")
	}

	return  nil
}
*/

func (s *Server) Start() {
	fmt.Printf("[zinx] server name: %s, listenner at tcp %s port %d \n",
		utils.GlobalObject.Name, utils.GlobalObject.Host, utils.GlobalObject.TcpPort)
	fmt.Printf("[zinx] version %s, MaxConn:%d, MaxPackageSize:%d, MaxWorkerTaskLen:%d \n",
		utils.GlobalObject.Version,
		utils.GlobalObject.MaxConn,
		utils.GlobalObject.MaxPackageSize,
		utils.GlobalObject.MaxWorkerTaskLen,
	)
	//fmt.Printf("[start] server listener at IP: %s, Port %d, is starting\n", s.IP, s.Port)

	go func() {
		//开启消息队列及worker工作池
		s.MsgHandler.StartWorkerPool()

		//获取一个tcp的addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error: ", err)
			return
		}

		//监听服务器的地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println(" listen ", s.IPVersion, "err", err)
			return
		}

		fmt.Println("start Zinx server succ", s.Name, "success, listening ... ")
		var cid uint32
		cid = 0

		//阻塞等等客户端链接，处理客户端业务（读写）
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("accept err", err)
				continue
			}

			//判断是否超过最大连接个数
			if s.ConnMgr.Len() >= utils.GlobalObject.MaxConn {
				//todo 给用户响应一个错误：超出最大连接数的错误包
				fmt.Println("too many connections! maxConn=", utils.GlobalObject.MaxConn)
				conn.Close()
				continue
			}

			//将处理新链接的业务方法 和 conn 进行绑定  得到我们的链接模块
			dealConn := NewConnection(s, conn, cid, s.MsgHandler)

			cid++

			//启动当前链接的业务处理
			go dealConn.Start()

			/*
				go func() {
					for {
						buf := make([]byte, 512)
						cnt, err := conn.Read(buf)
						if err != nil {
							fmt.Println("recv buf err ", err)
							continue
						}

						fmt.Printf("recv client buf %s, cnt=%d \n", buf, cnt)

						if _, err := conn.Write(buf[:cnt]); err !=nil {
							fmt.Println("write back buf err", err)
							continue
						}

					}
				}()
			*/

		}
	}()

}

//停止服务器 回收资源
func (s *Server) Stop() {
	fmt.Println("[stop] zinx server name", s.Name, " stop!")
	s.ConnMgr.CLearConn()
}

func (s *Server) Server() {
	s.Start()

	//todo

	select {}
}

func (s *Server) AddRouter(msgId uint32, router ziface.IRouter) {
	s.MsgHandler.AddRouter(msgId, router)
	fmt.Println("add router success ")
}

func (s *Server) GetConnMgr() ziface.IConnManager {
	return s.ConnMgr
}

func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:       utils.GlobalObject.Name,
		IPVersion:  "tcp4",
		IP:         utils.GlobalObject.Host,
		Port:       utils.GlobalObject.TcpPort,
		MsgHandler: NewMsgHandle(),
		ConnMgr:    NewConnManager(),
	}

	return s
}

//注册 OnConnStart 钩子函数的方法
func (s *Server) SetOnConnStart(hookFunc func(connection ziface.IConnection)) {
	s.OnConnStart = hookFunc
}

//注册 OnConnStop 钩子函数的方法
func (s *Server) SetOnConnStop(hookFunc func(connection ziface.IConnection)) {
	s.OnConnStop = hookFunc
}

//调用 OnConnStart 钩子函数的方法
func (s *Server) CallOnConnStart(conn ziface.IConnection) {
	if s.OnConnStart != nil {
		fmt.Println("----->> CallOnConnStart() ")
		s.OnConnStart(conn)
	}
}

//调用 OnConnStop 钩子函数的方法
func (s *Server) CallOnConnStop(conn ziface.IConnection) {
	if s.OnConnStop != nil {
		fmt.Println("----->> CallOnConnStop() ")
		s.OnConnStop(conn)
	}
}
