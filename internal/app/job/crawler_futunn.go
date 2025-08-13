package job

import (
	"gin-scaffold/internal/config"
)

func AddCrawlerFutunn() {
	//AJob
	//talentInfo
	name := "crawlerfutunn" // 必须全小写不能采用驼峰命名
	JobList = append(JobList, &Job{
		Status: config.Cfg.Job.JobStatus[name],
		Name:   name,
		Cron:   config.Cfg.Job.JobCron[name],
		Func:   addCrawlerFutunn,
	})
}

func addCrawlerFutunn() {
	//futunn.ConnectHtml()
}
