package thepaper

import (
	"context"
	"encoding/json"
	"fmt"
	"gin-scaffold/internal/config"
	"gin-scaffold/pkg/robot"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

// WorldNews 国际新闻
func WorldNews(ctx context.Context) {
	url := "https://www.thepaper.cn/channel_122908"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("priority", "u=0, i")
	req.Header.Add("sec-fetch-user", "?1")
	req.Header.Add("upgrade-insecure-requests", "1")
	req.Header.Add("Cookie", "ariaDefaultTheme=undefined; Hm_lvt_94a1e06bbce219d29285cee2e37d1d26=1755311313; Hm_lpvt_94a1e06bbce219d29285cee2e37d1d26=1755311313; HMACCOUNT=F38274C5C42977AB; blackAndWhiteMode=0; menuIds=[25949,143064,128409,26916,25950,122908,25951,154646,119908,136261,36079,119489,25952,25953,26161,-8,143022,143065,-21,-24,122153,143013,150010,-1]; tfstk=g4nZfsYnShKwVJfBfXEVa7230dZt9oRSomNbnxD0C5VGfG1m8Yhi5STTciz0TjUbsCnb0qlEaoOTSr4Hx-eigrpTGSVqGx06d3tSBAEYmSRWV36YXTRZuNq0sK23CoyGP3AYdAEYmBshOBf-BfLLRpa0mpr338b0mRbGKpy4USq0srfh-8V3ioc0oMV3n85cj5bgtpyYtSq0mjqHL-N3ioqmivc0pADcYWzMKGYGoMjRcPFoIWSc4TegS_kFlGjiYRzZmaFhmJiUQPPoQc91Nn2nADz7RtOzxY3IsJry0MqazVcaKvdV450SekVni1RQpVcZxroBHFHE05zojyWV2Yrgnj4qJQS7QlgUzcu9H6DiN5uuXx6y1vzr8ze3-t7gf4MSDzmMb_FQy-lzUf5N451YKjsmDV5cuP2LL79eLwO5P2S9k4GPkZU32JPWeLQAkP2LL79eLZQY-8eUNLpR.")
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Host", "www.thepaper.cn")
	req.Header.Add("Connection", "keep-alive")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// 解析HTML
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return
	}

	var scriptContent string
	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		// 检查当前script标签的id是否为__NEXT_DATA__
		if id, exists := s.Attr("id"); exists && id == "__NEXT_DATA__" {
			// 提取标签内的文本内容
			scriptContent = s.Text()
		}
	})
	if scriptContent == "" {
		return
	}
	thePaper := ThePaper{}
	json.Unmarshal([]byte(scriptContent), &thePaper)

	//msg := ""
	msgContent := [][]robot.ZhCnContent{}
	for _, content := range thePaper.Props.PageProps.Data.Data.List {
		msgContent = append(msgContent, []robot.ZhCnContent{
			{
				Tag:  "text",
				Text: fmt.Sprintf("[%s]%s", content.PubTime, content.Name),
			},
			{
				Tag:  "a",
				Text: "超链接",
				Href: "https://www.thepaper.cn/newsDetail_forward_" + content.ContId,
			},
		})
		//msg += fmt.Sprintf("[%s][%s](https://www.thepaper.cn/newsDetail_forward_%s)\n", content.PubTime, content.Name, content.ContId)
	}

	postMsg := robot.ZhCn{
		Title:   "澎湃新闻-国际新闻",
		Content: msgContent,
	}

	jsonData, _ := json.Marshal(postMsg)

	robot.SendFeishuRobotWithUrl(context.Background(), config.Cfg.FSRobot.NewsRobot, string(jsonData), robot.MsgTypePost)

}
func ChinaNews(ctx context.Context) {

}
