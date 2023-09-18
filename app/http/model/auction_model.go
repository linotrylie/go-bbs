package model

import (
	"fmt"
)

type Auction struct {
	changes                   map[string]interface{}
	Id                        int     `gorm:"primaryKey;column:id" json:"id"`
	Uid                       int     `gorm:"column:uid" json:"uid"`
	CostSingleGold            int     `gorm:"column:costsinglegold" json:"costsinglegold"`
	CostSingleCoin            int     `gorm:"column:costsinglecoin" json:"costsinglecoin"`
	CostSingleCopper          int     `gorm:"column:costsinglecopper" json:"costsinglecopper"`
	SaleSingleGold            int     `gorm:"column:salesinglegold" json:"salesinglegold"`
	SaleSingleCoin            int     `gorm:"column:salesinglecoin" json:"salesinglecoin"`
	SaleSingleCopper          int     `gorm:"column:salesinglecopper" json:"salesinglecopper"`
	Nums                      int     `gorm:"column:nums" json:"nums"`
	ItemName                  string  `gorm:"column:itemname" json:"itemname"`
	CostGold                  int     `gorm:"column:costgold" json:"costgold"`
	CostCoin                  int     `gorm:"column:costcoin" json:"costcoin"`
	CostCopper                int     `gorm:"column:costcopper" json:"costcopper"`
	SaleGold                  int     `gorm:"column:salegold" json:"salegold"`
	SaleCoin                  int     `gorm:"column:salecoin" json:"salecoin"`
	SaleCopper                int     `gorm:"column:salecopper" json:"salecopper"`
	HandlingChargeGold        int     `gorm:"column:handlingchargegold" json:"handlingchargegold"`
	HandlingChargeCoin        int     `gorm:"column:handlingchargecoin" json:"handlingchargecoin"`
	HandlingChargeCopper      int     `gorm:"column:handlingchargecopper" json:"handlingchargecopper"`
	ProfitGold                int     `gorm:"column:profitgold" json:"profitgold"`
	ProfitCoin                int     `gorm:"column:profitcoin" json:"profitcoin"`
	ProfitCopper              int     `gorm:"column:profitcopper" json:"profitcopper"`
	HandlingChargeTax         int     `gorm:"column:handlingchargetax" json:"handlingchargetax"`
	DailyIncomeTax            int     `gorm:"column:dailyincometax" json:"dailyincometax"`
	CreateTime                int64   `gorm:"column:create_time" json:"create_time"`
	HandlingChargeDailyGold   int     `gorm:"column:handlingchargedailygold" json:"handlingchargedailygold"`
	HandlingChargeDailyCoin   int     `gorm:"column:handlingchargedailycoin" json:"handlingchargedailycoin"`
	HandlingChargeDailyCopper int     `gorm:"column:handlingchargedailycopper" json:"handlingchargedailycopper"`
	Rmb                       float64 `gorm:"column:rmb" json:"rmb"`
	Ratio                     float64 `gorm:"column:ratio" json:"ratio"`
	DailyGold                 int     `gorm:"column:dailygold" json:"dailygold"`
	DailyCoin                 int     `gorm:"column:dailycoin" json:"dailycoin"`
	DailyCopper               int     `gorm:"column:dailycopper" json:"dailycopper"`
}

func (*Auction) TableName() string {
	return "bbs_auction"
}

// Location .
func (obj *Auction) Location() map[string]interface{} {
	return map[string]interface{}{"id": obj.Id}
}

// Redis Key .
func (obj *Auction) RedisKey() string {
	return obj.TableName() + "_" + fmt.Sprintf("%v", obj.Id)
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
func (obj *Auction) SetId(val int) *Auction {
	obj.Id += val
	obj.Update("id", obj.Id)
	return obj
}
func (obj *Auction) SetUid(val int) *Auction {
	obj.Uid = val
	obj.Update("uid", obj.Uid)
	return obj
}
func (obj *Auction) SetCostSingleGold(val int) *Auction {
	obj.CostSingleGold += val
	obj.Update("costsinglegold", obj.CostSingleGold)
	return obj
}
func (obj *Auction) SetCostSingleCoin(val int) *Auction {
	obj.CostSingleCoin += val
	obj.Update("costsinglecoin", obj.CostSingleCoin)
	return obj
}
func (obj *Auction) SetCostSingleCopper(val int) *Auction {
	obj.CostSingleCopper += val
	obj.Update("costsinglecopper", obj.CostSingleCopper)
	return obj
}
func (obj *Auction) SetSaleSingleGold(val int) *Auction {
	obj.SaleSingleGold += val
	obj.Update("salesinglegold", obj.SaleSingleGold)
	return obj
}
func (obj *Auction) SetSaleSingleCoin(val int) *Auction {
	obj.SaleSingleCoin += val
	obj.Update("salesinglecoin", obj.SaleSingleCoin)
	return obj
}
func (obj *Auction) SetSaleSingleCopper(val int) *Auction {
	obj.SaleSingleCopper += val
	obj.Update("salesinglecopper", obj.SaleSingleCopper)
	return obj
}
func (obj *Auction) SetNums(val int) *Auction {
	obj.Nums += val
	obj.Update("nums", obj.Nums)
	return obj
}
func (obj *Auction) SetItemName(val string) *Auction {
	obj.ItemName = val
	obj.Update("itemname", obj.ItemName)
	return obj
}
func (obj *Auction) SetCostGold(val int) *Auction {
	obj.CostGold += val
	obj.Update("costgold", obj.CostGold)
	return obj
}
func (obj *Auction) SetCostCoin(val int) *Auction {
	obj.CostCoin += val
	obj.Update("costcoin", obj.CostCoin)
	return obj
}
func (obj *Auction) SetCostCopper(val int) *Auction {
	obj.CostCopper += val
	obj.Update("costcopper", obj.CostCopper)
	return obj
}
func (obj *Auction) SetSaleGold(val int) *Auction {
	obj.SaleGold += val
	obj.Update("salegold", obj.SaleGold)
	return obj
}
func (obj *Auction) SetSaleCoin(val int) *Auction {
	obj.SaleCoin += val
	obj.Update("salecoin", obj.SaleCoin)
	return obj
}
func (obj *Auction) SetSaleCopper(val int) *Auction {
	obj.SaleCopper += val
	obj.Update("salecopper", obj.SaleCopper)
	return obj
}
func (obj *Auction) SetHandlingChargeGold(val int) *Auction {
	obj.HandlingChargeGold += val
	obj.Update("handlingchargegold", obj.HandlingChargeGold)
	return obj
}
func (obj *Auction) SetHandlingChargeCoin(val int) *Auction {
	obj.HandlingChargeCoin += val
	obj.Update("handlingchargecoin", obj.HandlingChargeCoin)
	return obj
}
func (obj *Auction) SetHandlingChargeCopper(val int) *Auction {
	obj.HandlingChargeCopper += val
	obj.Update("handlingchargecopper", obj.HandlingChargeCopper)
	return obj
}
func (obj *Auction) SetProfitGold(val int) *Auction {
	obj.ProfitGold += val
	obj.Update("profitgold", obj.ProfitGold)
	return obj
}
func (obj *Auction) SetProfitCoin(val int) *Auction {
	obj.ProfitCoin += val
	obj.Update("profitcoin", obj.ProfitCoin)
	return obj
}
func (obj *Auction) SetProfitCopper(val int) *Auction {
	obj.ProfitCopper += val
	obj.Update("profitcopper", obj.ProfitCopper)
	return obj
}
func (obj *Auction) SetHandlingChargeTax(val int) *Auction {
	obj.HandlingChargeTax += val
	obj.Update("handlingchargetax", obj.HandlingChargeTax)
	return obj
}
func (obj *Auction) SetDailyIncomeTax(val int) *Auction {
	obj.DailyIncomeTax += val
	obj.Update("dailyincometax", obj.DailyIncomeTax)
	return obj
}
func (obj *Auction) SetCreateTime(val int64) *Auction {
	obj.CreateTime += val
	obj.Update("create_time", obj.CreateTime)
	return obj
}
func (obj *Auction) SetHandlingChargeDailyGold(val int) *Auction {
	obj.HandlingChargeDailyGold += val
	obj.Update("handlingchargedailygold", obj.HandlingChargeDailyGold)
	return obj
}
func (obj *Auction) SetHandlingChargeDailyCoin(val int) *Auction {
	obj.HandlingChargeDailyCoin += val
	obj.Update("handlingchargedailycoin", obj.HandlingChargeDailyCoin)
	return obj
}
func (obj *Auction) SetHandlingChargeDailyCopper(val int) *Auction {
	obj.HandlingChargeDailyCopper += val
	obj.Update("handlingchargedailycopper", obj.HandlingChargeDailyCopper)
	return obj
}
func (obj *Auction) SetRmb(val float64) *Auction {
	obj.Rmb += val
	obj.Update("rmb", obj.Rmb)
	return obj
}
func (obj *Auction) SetRatio(val float64) *Auction {
	obj.Ratio += val
	obj.Update("ratio", obj.Ratio)
	return obj
}
func (obj *Auction) SetDailyGold(val int) *Auction {
	obj.DailyGold += val
	obj.Update("dailygold", obj.DailyGold)
	return obj
}
func (obj *Auction) SetDailyCoin(val int) *Auction {
	obj.DailyCoin += val
	obj.Update("dailycoin", obj.DailyCoin)
	return obj
}
func (obj *Auction) SetDailyCopper(val int) *Auction {
	obj.DailyCopper += val
	obj.Update("dailycopper", obj.DailyCopper)
	return obj
}
