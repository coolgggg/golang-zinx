package core

import "sync"

//当前游戏世界的总管理模块
type WorldManager struct {
	//当前世界地图的AOI规划管理器
	AoiMgr *AOIManager

	//当前在线的玩家集合
	Players map[int32]*Player

	//保护Players的互斥读写机制
	pLock sync.RWMutex
}

//提供一个对外的世界管理模块句柄
var WorldMgrObj *WorldManager

//初始化方法
func init() {
	WorldMgrObj = &WorldManager{
		//玩家集合
		Players: make(map[int32]*Player),
		//创建世界AOI地图规划
		AoiMgr: NewAOIManager(AOI_MIN_X, AOI_MAX_X, AOI_CNTS_X, AOI_MIN_Y, AOI_MAX_Y, AOI_CNTS_Y),
	}
}

//提供添加一个玩家的的功能，将玩家添加进玩家信息表Players
func (wm *WorldManager) AddPlayer(player *Player) {
	wm.pLock.Lock()
	wm.Players[player.Pid] = player
	wm.pLock.Unlock()

	wm.AoiMgr.AddToGridByPos(int(player.Pid), player.X, player.Z)
}

//从玩家信息表中移除一个玩家
func (wm *WorldManager) RemovePlayerByPid(pid int32) {
	//得到当前玩家
	player := wm.Players[pid]

	//从aoi mgr删除
	wm.AoiMgr.RemoveFromGridByPos(int(pid), player.X, player.Z)

	//从世界管理中删除
	wm.pLock.Lock()
	delete(wm.Players, pid)
	wm.pLock.Unlock()

}

//通过玩家ID 获取对应玩家信息
func (wm *WorldManager) GetPlayerByPid(pid int32) *Player {
	wm.pLock.RLock()
	defer wm.pLock.RUnlock()

	return wm.Players[pid]
}

//获取所有玩家的信息
func (wm *WorldManager) GetAllPlays() []*Player {
	wm.pLock.RLock()
	defer wm.pLock.RUnlock()

	players := make([]*Player, 0)

	for _, v := range wm.Players {
		players = append(players, v)
	}
	return players
}

//获取指定gid中的所有player信息
func (wm *WorldManager) GetPlayersByGid(gid int) []*Player {
	/*
		//通过gid获取 对应 格子中的所有pid
		pids := wm.AoiMgr.GetPidsByGid(gid)

		//通过pid找到对应的player对象
		players := make([]*Player, len(pids))

		wm.pLock.RLock()
		for _, pid := range pids {
			players = append(players, wm.Players(int32(pid)))
		}
		wm.pLock.RUnlock()

		return players
	*/
	return nil
}
