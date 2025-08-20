package common

const (
	StockTypeETF  = "etf"
	StockTypeCNSH = "cn-sh"
	StockTypeCNSZ = "cn-sz"

	WindowsName = "stock_push"
)

var StockList = map[string]map[string]string{
	//StockTypeETF: {
	//	"513120": "港股创新药ETF",
	//	"513090": "香港证券ETF",
	//	"513220": "中概互联ETF",
	//	"159887": "美国50ETF",
	//	"520550": "港股红利低波ETF",
	//	"159529": "普标消费ETF",
	//},
	//StockTypeCNSH: {
	//	"600872": "中炬高新",
	//	"601318": "中国平安",
	//},
	StockTypeCNSZ: {
		"000858": "五粮液",
		"002213": "大为股份",
		"300301": "ST长方",
	},
}

var StockValue = map[string]string{}

func InitStockValue() {
	StockValue = make(map[string]string)
}
