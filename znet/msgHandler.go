package znet

import (
	"fmt"
	"github.com/aceld/golang-zinx/utils"
	"github.com/aceld/golang-zinx/ziface"
	"strconv"
)

//消息处理模块的实现

type MsgHandle struct {
	//消息id和router对应关系的集合
	//存放每个MsgId锁对应的处理方法
	Apis map[uint32]ziface.IRouter
	//负责worker取任务的消息队列
	TaskQueue []chan ziface.IRequest
	//业务工作worker池的worker数量
	WorkerPoolSize uint32
}

//初始化、创建MsgHandle方法
func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis:           make(map[uint32]ziface.IRouter),
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize,
		TaskQueue:      make([]chan ziface.IRequest, utils.GlobalObject.WorkerPoolSize),
	}
}

//调度、执行对应的router消息处理方法
func (mh *MsgHandle) DoMsgHandler(request ziface.IRequest) {
	//从request中赵傲msg id
	handler, ok := mh.Apis[request.GetMsgId()]
	if !ok {
		fmt.Println("api msgId=", request.GetMsgId(), "is not found, need register")
	}

	//根据msg id调度对应的router业务
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}

//为消息添加具体的处理逻辑
func (mh *MsgHandle) AddRouter(msgID uint32, router ziface.IRouter) {
	//判断当前msg绑定的api处理方法是否已经存在
	if _, ok := mh.Apis[msgID]; ok {
		//id已经注册
		panic("repeat api, msg id = " + strconv.Itoa(int(msgID)))
	}

	//添加msg与api的绑定关系
	mh.Apis[msgID] = router
	fmt.Println("Add api msgId = ", msgID, "success!")
}

//启动一个worker工作池 (开启工作池的动作只能发生一次，一个框架只能有一个worker工作池)
func (mh *MsgHandle) StartWorkerPool() {
	//根据workerPoolSize 分别开启worker，每个worker用一个go承载
	for i := 0; i < int(mh.WorkerPoolSize); i++ {
		//一个worker被启动
		//当前的worker对应的channel消息队列开辟空间，第0个worker就用第0个channel
		mh.TaskQueue[i] = make(chan ziface.IRequest, utils.GlobalObject.MaxWorkerTaskLen)
		//启动当前worker，阻塞等待消息从channel传递进来
		go mh.StartOneWorker(i, mh.TaskQueue[i])
	}
}

//启动一个worker工作流程
func (mh *MsgHandle) StartOneWorker(workerId int, taskQueue chan ziface.IRequest) {
	fmt.Println("work id=", workerId, " is started")

	//不断阻塞等待对应消息队列的消息
	for {
		select {
		//如果有消息过来，出列的就是一个客户端的request，执行当前request说绑定的业务
		case request := <-taskQueue:
			mh.DoMsgHandler(request)
		}
	}
}

//将消息交给taskQueue，由worker进行处理
func (mh *MsgHandle) SendMsgToTaskQueue(request ziface.IRequest) {
	//将消息平均分配给不同的worker
	//根据客户端建立的connId来进行分配
	workId := request.GetConnection().GetConnID() % mh.WorkerPoolSize
	//workId := rand.Intn(10)
	fmt.Println("Add connId=", request.GetConnection().GetConnID(),
		" request MsgId=", request.GetMsgId(), " to workerId=", workId)

	//将消息发送给对应worker的taskQueue即可
	mh.TaskQueue[workId] <- request
}
