package crawler

import (
	"gin-scaffold/pkg/crawler/futunn"
	"time"
)

func InitCrawler() {
	//ctx := context.Background()

}

func CrawlerFutunn() {
	go func() {
		for {
			now := time.Now()                                                                                // 获取当前时间
			todayAt0900 := time.Date(now.Year(), now.Month(), now.Day(), 9, 20, 0, 0, now.Location()).Unix() // 构造当天9点的时间
			todayAt1130 := time.Date(now.Year(), now.Month(), now.Day(), 11, 30, 0, 0, now.Location()).Unix()
			todayAt1300 := time.Date(now.Year(), now.Month(), now.Day(), 13, 0, 0, 0, now.Location()).Unix()
			todayAt1500 := time.Date(now.Year(), now.Month(), now.Day(), 15, 10, 0, 0, now.Location()).Unix()
			nowUnix := time.Now().Unix()
			if (nowUnix >= todayAt0900 && nowUnix <= todayAt1130) || (nowUnix >= todayAt1300 && nowUnix <= todayAt1500) {
				futunn.ConnectHtml()
			} else {
				time.Sleep(5 * time.Minute)
			}
		}
	}()
}
