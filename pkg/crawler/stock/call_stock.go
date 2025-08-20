package stock

import (
	"gin-scaffold/internal/config"
	"gin-scaffold/pkg/crawler/stock/cfi"
	"gin-scaffold/pkg/crawler/stock/common"
	"gin-scaffold/pkg/crawler/stock/futunn"
	"time"
)

func CallStock() {
	stockList := config.Cfg.Stock

	for stockCode, stockName := range stockList.SZ {
		cfi.StockCFI(common.StockTypeCNSZ, stockCode, stockName)
		time.Sleep(10 * time.Second)
	}
	for stockCode, stockName := range stockList.SH {
		futunn.ConnectHtml(common.StockTypeCNSZ, stockCode, stockName)
		time.Sleep(10 * time.Second)
	}
	for stockCode, stockName := range stockList.ETF {
		futunn.ConnectHtml(common.StockTypeETF, stockCode, stockName)
		time.Sleep(10 * time.Second)
	}
}
