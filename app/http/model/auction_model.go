package model

type Auction struct {
	changes              map[string]interface{}
	Id                   int    `gorm:"primaryKey;column:id" json:"id"`
	Uid                  int    `gorm:"column:uid" json:"uid"`
	CostSingleGold       int    `gorm:"column:costsinglegold" json:"costsinglegold"`
	CostSingleCoin       int    `gorm:"column:costsinglecoin" json:"costsinglecoin"`
	CostSingleCopper     int    `gorm:"column:costsinglecopper" json:"costsinglecopper"`
	SaleSingleGold       int    `gorm:"column:salesinglegold" json:"salesinglegold"`
	SaleSingleCoin       int    `gorm:"column:salesinglecoin" json:"salesinglecoin"`
	SaleSingleCopper     int    `gorm:"column:salesinglecopper" json:"salesinglecopper"`
	Nums                 int    `gorm:"column:nums" json:"nums"`
	ItemName             string `gorm:"column:itemname" json:"itemname"`
	CostGold             int    `gorm:"column:costgold" json:"costgold"`
	CostCoin             int    `gorm:"column:costcoin" json:"costcoin"`
	CostCopper           int    `gorm:"column:costcopper" json:"costcopper"`
	SaleGold             int    `gorm:"column:salegold" json:"salegold"`
	SaleCoin             int    `gorm:"column:salecoin" json:"salecoin"`
	SaleCopper           int    `gorm:"column:salecopper" json:"salecopper"`
	HandlingChargeGold   int    `gorm:"column:handlingchargegold" json:"handlingchargegold"`
	HandlingChargeCoin   int    `gorm:"column:handlingchargecoin" json:"handlingchargecoin"`
	HandlingChargeCopper int    `gorm:"column:handlingchargecopper" json:"handlingchargecopper"`
	ProfitGold           int    `gorm:"column:profitgold" json:"profitgold"`
	ProfitCoin           int    `gorm:"column:profitcoin" json:"profitcoin"`
	ProfitCopper         int    `gorm:"column:profitcopper" json:"profitcopper"`
	HandlingChargeTax    int    `gorm:"column:handlingchargetax" json:"handlingchargetax"`
	DailyIncomeTax       int    `gorm:"column:dailyincometax" json:"dailyincometax"`
	CreateTime           string `gorm:"column:createtime" json:"createtime"`
	DailyGold            int    `gorm:"column:dailygold" json:"dailygold"`
	DailyCoin            int    `gorm:"column:dailycoin" json:"dailycoin"`
	DailyCopper          int    `gorm:"column:dailycopper" json:"dailycopper"`
}

func (*Auction) TableName() string {
	return "auction"
}

// Location .
func (obj *Auction) Location() map[string]interface{} {
	return map[string]interface{}{"Id": obj.Id}
}

// GetChanges .
func (obj *Auction) GetChanges() map[string]interface{} {
	if obj.changes == nil {
		return nil
	}
	result := make(map[string]interface{})
	for k, v := range obj.changes {
		result[k] = v
	}
	obj.changes = nil
	return result
}

// Update .
func (obj *Auction) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
