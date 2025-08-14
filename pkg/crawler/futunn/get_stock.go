package futunn

import (
	"context"
	"encoding/json"
	"fmt"
	"gin-scaffold/internal/config"
	"gin-scaffold/pkg/http_call"
	"gin-scaffold/pkg/logger"
	"gin-scaffold/pkg/utils"
	"github.com/PuerkitoBio/goquery"
	"math"
	"net/http"
	"regexp"
	"strings"
	"time"
)

const (
	StockTypeETF  = "etf"
	StockTypeCNSH = "cn-sh"
	StockTypeCNSZ = "cn-sz"
)

var stockList = map[string]map[string]string{
	StockTypeETF: {
		"513120": "港股创新药ETF",
		"513090": "香港证券ETF",
		"513220": "中概互联ETF",
		"159887": "美国50ETF",
		"520550": "港股红利低波ETF",
		"159529": "普标消费ETF",
	},
	StockTypeCNSH: {
		"600872": "中炬高新",
		"601318": "中国平安",
	},
	//StockTypeCNSZ: {
	//	"000858": "五粮液",
	//},
}

var stockValue = map[string]string{}

var futunnRatioRegex = regexp.MustCompile(`[+-]\d+\.\d+%`)
var futunnChangePriceRegex = regexp.MustCompile(`\d+\.\d+%`)

func ConnectHtml() {
	// 请求URL
	for stockType, stockTypeSon := range stockList {
		for stockCode, stockName := range stockTypeSon {
			time.Sleep(10 * time.Second)
			url := ""
			switch stockType {
			case StockTypeETF:
				url = fmt.Sprintf("https://www.futunn.com/etfs/%s-SH", stockCode)
			case StockTypeCNSH:
				url = fmt.Sprintf("https://www.futunn.com/stock/%s-SH", stockCode)
			case StockTypeCNSZ:
				//url = fmt.Sprintf("https://www.futunn.com/stock/%s-SZ", stockCode)
			}
			callFutunn(url, stockCode, stockName)
		}

	}
	// 打印响应状态码和内容
	//fmt.Printf("响应状态码: %d\n", resp.StatusCode)
	//fmt.Println("响应内容:")
	//fmt.Println(string(body))
}

func callFutunn(url string, stockCode, stockName string) {

	// 创建一个HTTP客户端
	client := &http.Client{}
	botUrl := config.Cfg.FSRobot.StockRobot

	// 创建请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("创建请求失败: %v\n", err)
		return
	}

	// 设置请求头
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("priority", "u=0, i")
	req.Header.Set("sec-ch-ua", `"Not;A=Brand";v="99", "Microsoft Edge";v="139", "Chromium";v="139"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-site", "none")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("upgrade-insecure-requests", "1")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/139.0.0.0 Safari/537.36 Edg/139.0.0.0")

	// 设置Cookie
	cookies := `csrfToken=cqQtDcK8j52pGMM5SLDMU7Qi; locale=zh-cn; locale.sig=ObiqV0BmZw7fEycdGJRoK-Q0Yeuop294gBeiHL1LqgQ; cipher_device_id=1755058562161456; device_id=1755058562161456; Hm_lvt_f3ecfeb354419b501942b6f9caf8d0db=1755058562; HMACCOUNT=F38274C5C42977AB; futu-csrf=Nk5ifXE3KlfKYFfIXSR6iD5E9as=; _gid=GA1.2.1348321045.1755058563; _gcl_au=1.1.1960921928.1755058564; sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%22ftv14FfjFXaKG0u9aodemHfGebRhebfgmuYMZvUSvxsSa%2FR7x8PSd0R0sg4cmCX4Ml24%22%2C%22first_id%22%3A%22198a1a46ec01ca3-03f1d019d88768c-4c657b58-2073600-198a1a46ec11dd9%22%2C%22props%22%3A%7B%22%24latest_traffic_source_type%22%3A%22%E7%9B%B4%E6%8E%A5%E6%B5%81%E9%87%8F%22%2C%22%24latest_search_keyword%22%3A%22%E6%9C%AA%E5%8F%96%E5%88%B0%E5%80%BC_%E7%9B%B4%E6%8E%A5%E6%89%93%E5%BC%80%22%2C%22%24latest_referrer%22%3A%22%22%7D%2C%22identities%22%3A%22eyIkaWRlbnRpdHlfY29va2llX2lkIjoiMTk4YTFhNDZlYzAxY2EzLTAzZjFkMDE5ZDg4NzY4Yy00YzY1N2I1OC0yMDczNjAwLTE5OGExYTQ2ZWMxMWRkOSIsIiRpZGVudGl0eV9sb2dpbl9pZCI6ImZ0djE0RmZqRlhhS0cwdTlhb2RlbUhmR2ViUmhlYmZnbXVZTVp2VVN2eHNTYS9SN3g4UFNkMFIwc2c0Y21DWDRNbDI0In0%3D%22%2C%22history_login_id%22%3A%7B%22name%22%3A%22%24identity_login_id%22%2C%22value%22%3A%22ftv14FfjFXaKG0u9aodemHfGebRhebfgmuYMZvUSvxsSa%2FR7x8PSd0R0sg4cmCX4Ml24%22%7D%7D; _ga_XECT8CPR37=GS2.1.s1755146175$o5$g1$t1755147371$j59$l0$h0; _ga=GA1.2.1839857836.1755058563; _gat_UA-71722593-3=1; Hm_lpvt_f3ecfeb354419b501942b6f9caf8d0db=1755147372; _ga_370Q8HQYD7=GS2.2.s1755146176$o5$g1$t1755147371$j60$l0$h0; ftreport-jssdk%40session={%22distinctId%22:%22ftv14FfjFXaKG0u9aodemHfGealQEu0CglA/sb6cjsBruZN7x8PSd0R0sg4cmCX4Ml24%22%2C%22firstId%22:%22ftv14FfjFXaKG0u9aodemHfGeQmesVqE+y2DezXEXaNRhGl7x8PSd0R0sg4cmCX4Ml24%22%2C%22latestReferrer%22:%22https://www.futunn.com/%22}; _ga_EJJJZFNPTW=GS2.1.s1755146161$o5$g1$t1755147376$j54$l0$h0`

	// 分割并设置每个Cookie
	for _, cookie := range strings.Split(cookies, "; ") {
		parts := strings.SplitN(cookie, "=", 2)
		if len(parts) == 2 {
			req.AddCookie(&http.Cookie{
				Name:  parts[0],
				Value: parts[1],
			})
		}
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("发送请求失败: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应内容
	//body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(body)
	//if err != nil {
	//	fmt.Printf("读取响应失败: %v\n", err)
	//	return
	//}

	// 解析HTML
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return
	}
	// 使用CSS选择器定位目标元素（根据页面结构）
	// 目标元素特征：class="change-ratio"的span标签
	//doc.Find("span.change-ratio").Each(func(i int, s *goquery.Selection) {
	//	text := s.Text()
	//	// 验证格式是否匹配
	//	if futunnRatioRegex.MatchString(text) {
	//		ratio = text
	//	}
	//})
	//doc.Find("span.change-price").Each(func(i int, s *goquery.Selection) {
	//	text := s.Text()
	//	// 验证格式是否匹配
	//	if futunnChangePriceRegex.MatchString(text) {
	//		changePrice = text
	//	}
	//})

	//doc.Find("ul.price-current").Each(func(i int, s *goquery.Selection) {
	//	text := s.Text()
	//	// 验证格式是否匹配
	//	//if futunnChangePriceRegex.MatchString(text) {
	//	currentPrice = text
	//	//}
	//})

	var scriptContent string
	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		if strings.Contains(text, "window.__INITIAL_STATE__=") {
			scriptContent = text
		}
	})

	if scriptContent == "" {
		return
	}

	// 提取JSON部分（截取=后面的内容，去除末尾分号）
	re := regexp.MustCompile(`window\.__INITIAL_STATE__=(.*?);?\s*$`)
	matches := re.FindStringSubmatch(scriptContent)
	if len(matches) < 2 {
		return
	}
	jsonStr := matches[1]
	splits := strings.Split(jsonStr, ";")
	if len(splits) == 0 {
		return
	}
	jsonStr = splits[0]

	// 解析JSON
	var state Futunn
	if err = json.Unmarshal([]byte(jsonStr), &state); err != nil {
		return
	}

	ratio := state.StockInfo.ChangeRatio
	changePrice := state.StockInfo.Change
	currentPrice := state.StockInfo.Price
	if ratio == "" {
		fmt.Println("未找到涨跌幅数据")
		return
	}
	if val, ok := stockValue[stockCode]; ok {
		subValue := utils.ExtractPercentageNumber(val) - utils.ExtractPercentageNumber(ratio)
		msg := fmt.Sprintf("[%s]%s(%s) 当前涨幅:%s  当前价:%s", ratio, stockName, stockCode, changePrice, currentPrice)
		if subValue > 0.5 {
			msg = fmt.Sprintf("↓[%s]%s(%s) 当前涨幅:%s  当前价:%s", ratio, stockName, stockCode, changePrice, currentPrice)
		}
		if subValue < -0.5 {
			msg = fmt.Sprintf("↑[%s]%s(%s) 当前涨幅:%s  当前价:%s", ratio, stockName, stockCode, changePrice, currentPrice)
		}
		logger.Logger.Info(msg)
		if math.Abs(subValue) >= 0.5 {
			http_call.SendFeishuRobotWithUrl(context.Background(), botUrl, msg)
			stockValue[stockCode] = ratio
		}
		return
	}

	stockValue[stockCode] = ratio
	msg := fmt.Sprintf("[%s]%s(%s) 当前涨幅:%s  当前价:%s", ratio, stockName, stockCode, changePrice, currentPrice)
	http_call.SendFeishuRobotWithUrl(context.Background(), botUrl, msg)
	logger.Logger.Info(msg)
}
