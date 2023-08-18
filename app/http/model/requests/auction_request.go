package requests

type AuctionRequest struct {
	Id                   int    ` json:"id"`
	Uid                  int    ` json:"uid"`
	CostSingleGold       int    ` json:"costsinglegold"`
	CostSingleCoin       int    ` json:"costsinglecoin"`
	CostSingleCopper     int    ` json:"costsinglecopper"`
	SaleSingleGold       int    ` json:"salesinglegold"`
	SaleSingleCoin       int    ` json:"salesinglecoin"`
	SaleSingleCopper     int    ` json:"salesinglecopper"`
	Nums                 int    ` json:"nums"`
	ItemName             string ` json:"itemname"`
	CostGold             int    ` json:"costgold"`
	CostCoin             int    ` json:"costcoin"`
	CostCopper           int    ` json:"costcopper"`
	SaleGold             int    ` json:"salegold"`
	SaleCoin             int    ` json:"salecoin"`
	SaleCopper           int    ` json:"salecopper"`
	HandlingChargeGold   int    ` json:"handlingchargegold"`
	HandlingChargeCoin   int    ` json:"handlingchargecoin"`
	HandlingChargeCopper int    ` json:"handlingchargecopper"`
	ProfitGold           int    ` json:"profitgold"`
	ProfitCoin           int    ` json:"profitcoin"`
	ProfitCopper         int    ` json:"profitcopper"`
	HandlingChargeTax    int    ` json:"handlingchargetax"`
	DailyIncomeTax       int    ` json:"dailyincometax"`
	CreateTime           string ` json:"createtime"`
	DailyGold            int    ` json:"dailygold"`
	DailyCoin            int    ` json:"dailycoin"`
	DailyCopper          int    ` json:"dailycopper"`
}
