package cfg_go
/**
由 Cfg_Shop.xlsx 商店 excel文件生成 ...
author:yh 
*/
type Cfg_Shop struct{
	Shelves map[int32]*Cfg_Shop_Shelves_Item
}
type Cfg_Shop_Shelves_Item struct {
   /* 货架ID */
	Id int32 `json:"Id"`
   /* 货架名称 */
	Title string `json:"Title"`
   /* 商品ID */
	Com []int32 `json:"Com"`
   /* 商品ID */
	CommodityWeight [][]int32 `json:"CommodityWeight"`
   /* 是否刷新(1、刷新;2、不刷新) */
	Rush string `json:"Rush"`
   /* 刷新倒计时 */
	RushTime string `json:"RushTime"`
   /* 商品ID */
	ssss int32 `json:"ssss"`
}
func (t *Cfg_Shop_Shelves_Item) Init(Id int32,Title string,Com []int32,CommodityWeight [][]int32,Rush string,RushTime string,ssss int32) {
	  t.Id = Id
	  t.Title = Title
	  t.Com = Com
	  t.CommodityWeight = CommodityWeight
	  t.Rush = Rush
	  t.RushTime = RushTime
	  t.ssss = ssss
	}
func (t *Cfg_Shop_Shelves_Item) Clone() *Cfg_Shop_Shelves_Item {
	return &Cfg_Shop_Shelves_Item{
		Id:t.Id,
		Title:t.Title,
		Com:t.Com,
		CommodityWeight:t.CommodityWeight,
		Rush:t.Rush,
		RushTime:t.RushTime,
		ssss:t.ssss,
	}
}

