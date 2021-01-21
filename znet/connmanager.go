package znet

import (
	"errors"
	"fmt"
	"golang-zinx/ziface"
	"sync"
)

//连接管理模块
type ConnManager struct {
	connections map[uint32]ziface.IConnection //管理的连接集合
	connLock    sync.RWMutex                  //互斥锁
}

//创建当前连接的方法
func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[uint32]ziface.IConnection),
	}
}

//添加连接
func (connMgr *ConnManager) Add(conn ziface.IConnection) {
	//保护共享资源map，加写锁
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()

	//将conn加入到ConnManager中
	connMgr.connections[conn.GetConnID()] = conn
	fmt.Println("connId=", conn.GetConnID(), " add to ConnManager successful, conn num = ", connMgr.Len())
}

//删除连接
func (connMgr *ConnManager) Remove(conn ziface.IConnection) {
	//保护共享资源map，加写锁
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()

	//删除连接信息
	delete(connMgr.connections, conn.GetConnID())
	fmt.Println("connId=", conn.GetConnID(), " remove from ConnManager successful, conn num = ", connMgr.Len())
}

//根据connId获取连接
func (connMgr *ConnManager) Get(connId uint32) (ziface.IConnection, error) {
	//保护共享资源map，加读锁
	connMgr.connLock.RLock()
	defer connMgr.connLock.RUnlock()

	if conn, ok := connMgr.connections[connId]; ok {
		return conn, nil
	} else {
		return nil, errors.New("connection not found ")
	}
}

//得到当前连接数
func (connMgr *ConnManager) Len() int {
	return len(connMgr.connections)
}

//清除并终止所有连接
func (connMgr *ConnManager) CLearConn() {
	//保护共享资源map，加写锁
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()

	//删除conn并停止conn的工作
	for connId, conn := range connMgr.connections {
		//停止
		conn.Stop()
		//删除
		delete(connMgr.connections, connId)
	}
	fmt.Println("clear all connections succ, conn num:", connMgr.Len())
}
