package stock

import (
	"gin-scaffold/pkg/crawler/stock/futunn"
	"gin-scaffold/pkg/logger"
	"time"
)

func InitCrawler() {
	//ctx := context.Background()

}

func CrawlerFutunn() {
	go func() {
		for {
			futunn.ConnectHtml()
			now := time.Now() // 获取当前时间
			if now.Weekday() == time.Saturday || now.Weekday() == time.Sunday {
				continue
			}
			todayAt0900 := time.Date(now.Year(), now.Month(), now.Day(), 9, 20, 0, 0, now.Location()).Unix() // 构造当天9点的时间
			todayAt1130 := time.Date(now.Year(), now.Month(), now.Day(), 11, 30, 0, 0, now.Location()).Unix()
			todayAt1300 := time.Date(now.Year(), now.Month(), now.Day(), 13, 0, 0, 0, now.Location()).Unix()
			todayAt1500 := time.Date(now.Year(), now.Month(), now.Day(), 15, 10, 0, 0, now.Location()).Unix()
			nowUnix := time.Now().Unix()
			if (nowUnix >= todayAt0900 && nowUnix <= todayAt1130) || (nowUnix >= todayAt1300 && nowUnix <= todayAt1500) {
				futunn.ConnectHtml()
			} else {
				futunn.InitStockValue()
				logger.Logger.Info("休眠5分钟 - 非交易时间")
				time.Sleep(5 * time.Minute)
			}
		}
	}()
}
