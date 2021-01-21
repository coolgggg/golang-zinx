package znet

import (
	"errors"
	"fmt"
	"golang-zinx/utils"
	"golang-zinx/ziface"
	"io"
	"net"
	"sync"
)

//链接模块
type Connection struct {
	//当前conn创建时 隶属于那个server
	TcpServer ziface.IServer

	//当前链接的 socket tcp套接字
	Conn *net.TCPConn

	//当前链接的id
	ConnID uint32

	//当前链接的状态
	isClosed bool

	//当前链接所绑定的处理业务方法api
	//handleAPI ziface.HandleFunc

	//告知当前链接已经退出/停止的channel (由reader告知writer退出)
	ExitChan chan bool

	//无缓存的管道，用于读写goroutine直接的消息通信
	msgChan chan []byte

	//该链接处理的方法router
	//Router ziface.IRouter

	//消息的管理msgId和对应的业务处理api关系
	MsgHandler ziface.IMsgHandle

	//链接属性集合
	property map[string]interface{}

	//包含链接属性的锁
	propertyLock sync.RWMutex
}

//初始化链接模块的方法
func NewConnection(server ziface.IServer, conn *net.TCPConn, connID uint32, msgHandler ziface.IMsgHandle) *Connection {
	c := &Connection{
		TcpServer:  server,
		Conn:       conn,
		ConnID:     connID,
		MsgHandler: msgHandler,
		isClosed:   false,
		msgChan:    make(chan []byte),
		ExitChan:   make(chan bool, 1),
		property:   make(map[string]interface{}),
	}

	//将conn加入到 connManager中
	c.TcpServer.GetConnMgr().Add(c)

	return c
}

func (c *Connection) StartReader() {
	fmt.Println("[read goroutine is running]")
	defer fmt.Println("[reader is exit] conn id =", c.ConnID, " remote addr=", c.RemoteAddr().String())
	defer c.Stop()

	for {
		/*
			//读取客户端的数据到buf中，目前最大就是512字节
			buf := make([]byte, utils.GlobalObject.MaxPackageSize)
			_, err := c.Conn.Read(buf)
			if err != nil {
				fmt.Println("recv buf err", err)
				//continue
				break
			}
		*/

		//创建一个拆包解包的对象
		dp := NewDataPack()

		//读取客户端的msg head 二进制流 8个字节
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
			fmt.Println("read msg head error", err)
			break
		}

		//拆包，得到客户端的 msg id 和 msg data len 放到msg消息中
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("unpack error ", err)
			break
		}

		//根据data len，再次读取 data，放在msg.DAta中
		var data []byte
		if msg.GetMsgLen() > 0 {
			data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("read msg data error", err)
				break
			}
		}
		msg.SetData(data)

		//得到当前conn数据的request请求数据
		req := Request{
			conn: c,
			msg:  msg,
		}

		/*
			//执行注册的路由方法
			go func(request ziface.IRequest) {
				//从路由中，找到注册绑定的con对应的router调用
				c.Router.PreHandle(request)
				c.Router.Handle(request)
				c.Router.PostHandle(request)
			}(&req)
		*/

		if utils.GlobalObject.WorkerPoolSize > 0 {
			//已经开启了工作池机制，将消息发送给worker工作池即可
			c.MsgHandler.SendMsgToTaskQueue(&req)
		} else {
			//根据绑定好的MsgId找到对应的处理api业务 执行
			go c.MsgHandler.DoMsgHandler(&req)
		}

		/*
			fmt.Printf("recv client buf %s, cnt=%d \n", buf, cnt)
			//调用当前链接锁绑定的的HandleAPI
			if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
				fmt.Println("conn id ", c.ConnID, "handle is err", err)
				break
			}
		*/
	}
}

//写消息的goroutine，专门发送给客户端消息的模块
func (c *Connection) StartWriter() {
	fmt.Println("[Write goroutine is running]")
	defer fmt.Println("[conn Writer exit!]", c.RemoteAddr().String())

	//不断的阻塞等待channel的消息，进行写给客户端
	for {
		select {
		case data := <-c.msgChan:
			//有数据要写给客户端
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("send data err", err)
				return
			}
		case <-c.ExitChan:
			//代表reader已经退出，此时writer也要退出
			return
		}
	}
}

//启动链接 让当前的链接开始工作
func (c *Connection) Start() {
	fmt.Println("connect start .. conn id=", c.ConnID)

	//启动当前链接写业务数据
	go c.StartReader()

	//启动当前链接写业务数据
	go c.StartWriter()

	//调用开发者传递进来的 创建连接之后需要调用的处理业务，执行对应的hook函数
	c.TcpServer.CallOnConnStart(c)
}

//停止链接 结束当前链接的工作
func (c *Connection) Stop() {
	fmt.Println("conn stop() .. conn id = ", c.ConnID)

	if c.isClosed == true {
		return
	}

	c.isClosed = true

	//调用开发者注册的 销毁连接之前 需要执行的hook函数
	c.TcpServer.CallOnConnStop(c)

	//关闭socket链接
	c.Conn.Close()

	//告知writer关闭
	c.ExitChan <- true

	//将当前连接从conn中摘除掉
	c.TcpServer.GetConnMgr().Remove(c)

	//回收资源
	close(c.ExitChan)
	close(c.msgChan)
}

//获取当前链接的绑定 socket conn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

//获取当前链接模块的ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

//获取远程客户端的 tcp状态 ip port
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

//发送数据，将数据发送给远程的客户端
//提供一个send msg 方法，将我们要发送给客户端的数据，先进行封包 在发送
func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed == true {
		return errors.New("connection close when send msg")
	}

	//将data进行封包 MsgDataLen|MsgId|Data
	dp := NewDataPack()

	//MsgDataLen|MsgID|Data
	binaryMsg, err := dp.Pack(NewMesPackage(msgId, data))
	if err != nil {
		fmt.Println("msg pack err, id:", msgId)
		return errors.New("pack error msg")
	}

	//将数据发送给客户端
	/*
		if _, err := c.Conn.Write(binaryMsg); err != nil {
			fmt.Println("write msg id", msgId, "err:", err)
			return errors.New("coon write err")
		}
	*/
	//将数据发送给chan
	c.msgChan <- binaryMsg

	return nil
}

//设置链接属性
func (c *Connection) SetProperty(key string, value interface{}) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	c.property[key] = value
}

//获取链接属性
func (c *Connection) GetProperty(key string) (interface{}, error) {
	c.propertyLock.RLock()
	defer c.propertyLock.RUnlock()
	//读取属性
	if value, ok := c.property[key]; ok {
		return value, nil
	} else {
		return nil, errors.New("no property found")
	}
}

//移除链接属性
func (c *Connection) RemoveProperty(key string) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	delete(c.property, key)
}
