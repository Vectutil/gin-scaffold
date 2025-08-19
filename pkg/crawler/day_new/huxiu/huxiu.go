package huxiu

import (
	"context"
	"encoding/json"
	"fmt"
	"gin-scaffold/internal/config"
	"gin-scaffold/pkg/robot"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var huxiuChan = map[string]string{
	"105": "前沿科技",
	"103": "商业消费",
	"106": "社会文化",
	"115": "金融财经",
	"107": "国际热点",
	"121": "3C热点",
	"123": "其他",
}

func HuxiuNews(ctx context.Context) {
	for newType, newTypeName := range huxiuChan {
		huxiuNews(ctx, newType, newTypeName)
		time.Sleep(10 * time.Second)
	}
}

func huxiuNews(ctx context.Context, newType, newTypeName string) {
	// 请求URL
	url := "https://www.huxiu.com/channel/" + newType + ".html"

	// 创建HTTP客户端
	client := &http.Client{}

	// 创建请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("创建请求失败: %v\n", err)
		return
	}

	// 设置请求头
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("cache-control", "max-age=0")
	req.Header.Set("priority", "u=0, i")
	req.Header.Set("referer", "https://www.huxiu.com/")
	req.Header.Set("sec-ch-ua", `"Not;A=Brand";v="99", "Microsoft Edge";v="139", "Chromium";v="139"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("upgrade-insecure-requests", "1")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/139.0.0.0 Safari/537.36 Edg/139.0.0.0")

	// 设置Cookie
	cookies := []*http.Cookie{
		{Name: "_ga", Value: "GA1.1.1752943396.1755219320"},
		{Name: "huxiu_analyzer_wcy_id", Value: "2zrr22kcjt5oncju9xgj"},
		{Name: "Hm_lvt_502e601588875750790bbe57346e972b", Value: "1755219320,1755481660"},
		{Name: "HMACCOUNT", Value: "F38274C5C42977AB"},
		{Name: "_ga_QRKQ6T69K7", Value: "GS2.1.s1755481658$o2$g1$t1755497549$j60$l0$h0"},
		{Name: "Hm_lpvt_502e601588875750790bbe57346e972b", Value: "1755497549"},
	}
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("发送请求失败: %v\n", err)
		return
	}
	defer resp.Body.Close()
	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	// 解析 HTML 内容为 goquery.Document
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
	if err != nil {
		log.Fatalf("解析 HTML 失败: %v", err)
	}
	cardMsg := robot.FeishuWebhookRequestCard{}
	cardMsg.Elements = append(cardMsg.Elements, robot.Elements{
		Tag: "div",
		Text: robot.ElementsText{
			Content: fmt.Sprintf("**虎嗅新闻 - %s**", newTypeName),
			Tag:     "lark_md",
		},
	})
	// 定位所有 article-item-wrap 节点（此处假设只有一个目标节点）
	doc.Find(".article-item-wrap").Each(func(i int, s *goquery.Selection) {
		// 提取链接（取第一个 a 标签的 href 属性，或标题所在 a 标签的 href）
		link, exists := s.Find("a").First().Attr("href")
		if !exists {
			link = "未找到链接"
		}

		// 提取标题（h3.channel-title 的文本）
		title := s.Find("h3.channel-title").Text()
		if title == "" {
			title = "未找到标题"
			
		}

		// 提取时间（div.bottom-line__time 的文本）
		time := s.Find("div.bottom-line__time").Text()
		if time == "" {
			time = "未找到时间"
		}

		cardMsg.Elements = append(cardMsg.Elements, robot.Elements{
			Tag: "div",
			Text: robot.ElementsText{
				Content: fmt.Sprintf("<font color='grey'>(%s)<font> [**%s**](%s)", time, title, link),
				Tag:     "lark_md",
			},
		})
	})
	jsonData, _ := json.Marshal(cardMsg)
	robot.SendFeishuRobotWithUrl(context.Background(), config.Cfg.FSRobot.NewsRobot, string(jsonData), robot.MsgTypeInteractive)
}
