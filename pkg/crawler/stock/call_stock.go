package stock

import (
	"gin-scaffold/pkg/crawler/stock/cfi"
	"gin-scaffold/pkg/crawler/stock/common"
	"gin-scaffold/pkg/crawler/stock/futunn"
	"time"
)

func CallStock() {
	for stockType, stockTypeList := range common.StockList {
		for stockCode, stockName := range stockTypeList {
			switch stockType {
			case common.StockTypeETF:
				futunn.ConnectHtml(stockType, stockCode, stockName)
			case common.StockTypeCNSH:
				futunn.ConnectHtml(stockType, stockCode, stockName)
			case common.StockTypeCNSZ:
				cfi.StockCFI(stockType, stockCode, stockName)
			}

			time.Sleep(10 * time.Second)
		}

	}
}
