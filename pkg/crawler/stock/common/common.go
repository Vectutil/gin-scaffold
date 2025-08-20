package common

const (
	StockTypeETF  = "etf"
	StockTypeCNSH = "cn-sh"
	StockTypeCNSZ = "cn-sz"

	WindowsName = "stock_push"
)

var StockValue = map[string]string{}

func InitStockValue() {
	StockValue = make(map[string]string)
}
