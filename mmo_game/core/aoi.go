package core

import "fmt"

const (
	AOI_MIN_X  int = 85
	AOI_MAX_X  int = 410
	AOI_CNTS_X int = 10
	AOI_MIN_Y  int = 75
	AOI_MAX_Y  int = 400
	AOI_CNTS_Y int = 20
)

/*
   AOI管理模块
*/
type AOIManager struct {
	MinX  int           //区域左边界坐标
	MaxX  int           //区域右边界坐标
	CntsX int           //x方向格子的数量
	MinY  int           //区域上边界坐标
	MaxY  int           //区域下边界坐标
	CntsY int           //y方向的格子数量
	grids map[int]*Grid //当前区域中都有哪些格子，key=格子ID， value=格子对象
}

/*
	初始化一个AOI区域
*/
func NewAOIManager(minX, maxX, cntsX, minY, maxY, cntsY int) *AOIManager {
	aoiMgr := &AOIManager{
		MinX:  minX,
		MaxX:  maxX,
		CntsX: cntsX,
		MinY:  minY,
		MaxY:  maxY,
		CntsY: cntsY,
		grids: make(map[int]*Grid),
	}

	//给AOI初始化区域中所有的格子
	for y := 0; y < cntsY; y++ {
		for x := 0; x < cntsX; x++ {
			//计算格子ID
			//格子编号：id = idy *nx + idx  (利用格子坐标得到格子编号)
			gid := y*cntsX + x

			//初始化一个格子放在AOI中的map里，key是当前格子的ID
			aoiMgr.grids[gid] = NewGrid(gid,
				aoiMgr.MinX+x*aoiMgr.gridWidth(),
				aoiMgr.MinX+(x+1)*aoiMgr.gridWidth(),
				aoiMgr.MinY+y*aoiMgr.gridLength(),
				aoiMgr.MinY+(y+1)*aoiMgr.gridLength())
		}
	}

	return aoiMgr
}

//得到每个格子在x轴方向的宽度
func (m *AOIManager) gridWidth() int {
	return (m.MaxX - m.MinX) / m.CntsX
}

//得到每个格子在x轴方向的长度
func (m *AOIManager) gridLength() int {
	return (m.MaxY - m.MinY) / m.CntsY
}

//打印信息方法
func (m *AOIManager) String() string {
	s := fmt.Sprintf("AOIManagr:\nminX:%d, maxX:%d, cntsX:%d, minY:%d, maxY:%d, cntsY:%d\n Grids in AOI Manager:\n",
		m.MinX, m.MaxX, m.CntsX, m.MinY, m.MaxY, m.CntsY)
	for _, grid := range m.grids {
		s += fmt.Sprintln(grid)
	}

	return s
}

//根据格子GID得到周边九宫格格子的ID集合
func (m *AOIManager) GetSurroundGridsByGid(gID int) (grids []*Grid) {
	//判断gID是否在AOIManager中
	if _, ok := m.grids[gID]; !ok {
		return
	}

	//将当前gid本身加入九宫格切片中
	grids = append(grids, m.grids[gID]) //假设是8

	//需要gID的左边是否有格子，右边是否有格子
	//需要通过gID得到当前格子的x轴编号 --idx = id % nx
	idx := gID % m.CntsX // 8%5  = 3

	//判断idx编号是否左边还有格子,如果有 放在gidsX集合中
	if idx > 0 {
		grids = append(grids, m.grids[gID-1]) //7
	}

	//判断idx编号是否右边还有格子
	if idx < m.CntsX-1 {
		grids = append(grids, m.grids[gID+1]) //9
	}

	//将x轴当前的格子都取出，进行遍历，再分别得到每个格子的上下是否有格子
	//得到当前x轴格子的ID集合
	gidsX := make([]int, 0, len(grids))
	for _, v := range grids {
		gidsX = append(gidsX, v.GID)
	}

	//遍历gidsX集合中每个格子的gid
	for _, v := range gidsX {
		//得到当前格子id的Y轴编号 idy = id / ny
		idy := v / m.CntsY
		//gid上边是否还有格子
		if idy > 0 {
			grids = append(grids, m.grids[v-m.CntsX])
		}

		//gid下边是否还有格子
		if idy < m.CntsX-1 {
			grids = append(grids, m.grids[v+m.CntsX])
		}
	}

	return
}

//通过横纵轴坐标得到GID
func (m *AOIManager) GetGidByPos(x, y float32) int {
	idx := (int(x) - m.MinX) / m.gridWidth()
	idy := (int(y) - m.MinY) / m.gridLength()

	return idy*m.CntsX + idx
}

//通过横纵坐标得到全部的PlayerIDs
func (m *AOIManager) GetPidsByPos(x, y float32) (playerIDs []int) {
	//得到当前玩家的GID格子id
	gID := m.GetGidByPos(x, y)

	//通过GID得到九宫格信息
	grids := m.GetSurroundGridsByGid(gID)

	//将九宫格的信息立的全部 PLAYER 累加到 playerIDs
	for _, grid := range grids {
		playerIDs = append(playerIDs, grid.GetPlayerIDs()...)
		fmt.Println("===> grid ID: %d, pids: $v ===", grid.GID, grid.GetPlayerIDs())
	}

	return
}

//添加一个 PlayerID 到格子中
func (m *AOIManager) AddPidToGrid(pID, gID int) {
	m.grids[gID].Add(pID)
}

//移除一个格子中的 PlayerID
func (m *AOIManager) RemovePidFromGrid(pID, gID int) {
	m.grids[gID].Remove(pID)
}

//通过GID获取全部的 PlayerID
func (m *AOIManager) GetPidsByGid(gID int) (playerIDs []int) {
	playerIDs = m.grids[gID].GetPlayerIDs()
	return
}

//通过坐标将 Player 添加到一个格子中
func (m *AOIManager) AddToGridByPos(pID int, x, y float32) {
	gID := m.GetGidByPos(x, y)
	grid := m.grids[gID]
	grid.Add(pID)
}

//通过坐标把一个 Player 从一个格子中删除
func (m *AOIManager) RemoveFromGridByPos(pID int, x, y float32) {
	gID := m.GetGidByPos(x, y)
	grid := m.grids[gID]
	grid.Remove(pID)
}
