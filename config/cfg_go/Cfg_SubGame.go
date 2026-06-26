package cfg_go
/**
由 Cfg_SubGame.xlsx 子游戏 excel文件生成 ...
author:yh 
*/
type Cfg_SubGame struct{
	Config map[int32]*Cfg_SubGame_Config_Item
	Item map[int32]*Cfg_SubGame_Item_Item
}
type Cfg_SubGame_Config_Item struct {
   /* 子游戏Id, */
	Id int32 `json:"Id"`
   /* 房间类型 */
	SubGameId int32 `json:"SubGameId"`
   /* 子游戏名字 */
	Name string `json:"Name"`
   /* 房间等级 */
	RoomLevel int32 `json:"RoomLevel"`
   /* 最小下注 */
	MiniBet int32 `json:"MiniBet"`
   /* 最大下注 */
	MaxBet int32 `json:"MaxBet"`
}
func (t *Cfg_SubGame_Config_Item) Init(Id int32,SubGameId int32,Name string,RoomLevel int32,MiniBet int32,MaxBet int32) {
	  t.Id = Id
	  t.SubGameId = SubGameId
	  t.Name = Name
	  t.RoomLevel = RoomLevel
	  t.MiniBet = MiniBet
	  t.MaxBet = MaxBet
	}
func (t *Cfg_SubGame_Config_Item) Clone() *Cfg_SubGame_Config_Item {
	return &Cfg_SubGame_Config_Item{
		Id:t.Id,
		SubGameId:t.SubGameId,
		Name:t.Name,
		RoomLevel:t.RoomLevel,
		MiniBet:t.MiniBet,
		MaxBet:t.MaxBet,
	}
}

type Cfg_SubGame_Item_Item struct {
   /* 房间类型 */
	SubGameId int32 `json:"SubGameId"`
   /* 子游戏名字 */
	Name string `json:"Name"`
   /* 资产名字 */
	AssetName string `json:"AssetName"`
   /* 图标 */
	Icon string `json:"Icon"`
}
func (t *Cfg_SubGame_Item_Item) Init(SubGameId int32,Name string,AssetName string,Icon string) {
	  t.SubGameId = SubGameId
	  t.Name = Name
	  t.AssetName = AssetName
	  t.Icon = Icon
	}
func (t *Cfg_SubGame_Item_Item) Clone() *Cfg_SubGame_Item_Item {
	return &Cfg_SubGame_Item_Item{
		SubGameId:t.SubGameId,
		Name:t.Name,
		AssetName:t.AssetName,
		Icon:t.Icon,
	}
}

