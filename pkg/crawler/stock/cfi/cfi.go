package cfi

import (
	"fmt"
	"gin-scaffold/pkg/crawler/stock/common"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

func StockCFI(stockType, stockCode, stockName string) {
	url := ""
	switch stockType {
	case common.StockTypeCNSZ:
		url = fmt.Sprintf("https://quote.cfi.cn/quote_%s.html", stockCode)
	}
	stockCFI(url, stockCode, stockName)
}

func stockCFI(url, stockCode, stockName string) {
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Add("Cache-Control", "max-age=0")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Sec-Fetch-Dest", "document")
	req.Header.Add("Sec-Fetch-Mode", "navigate")
	req.Header.Add("Sec-Fetch-Site", "same-site")
	req.Header.Add("Sec-Fetch-User", "?1")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/139.0.0.0 Safari/537.36 Edg/139.0.0.0")
	req.Header.Add("sec-ch-ua", "\"Not;A=Brand\";v=\"99\", \"Microsoft Edge\";v=\"139\", \"Chromium\";v=\"139\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Add("Host", "quote.cfi.cn")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	//all, _ := io.ReadAll(res.Body)
	//fmt.Println(string(all))

	//os.ReadFile()

	// 解析 HTML
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	price := ""
	change := ""
	doc.Find("#last").Each(func(i int, s *goquery.Selection) {
		price = strings.TrimSuffix(s.Text(), "↑") // 去除箭头
	})
	doc.Find("#chg").Each(func(i int, s *goquery.Selection) {
		change = s.Text()
	})
	common.CallRobot(price, change[0:4], change[4:], stockCode, stockName)
}
