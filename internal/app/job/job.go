package job

import (
	"gin-scaffold/pkg/logger"
	"github.com/robfig/cron/v3"
)

var (
	JobList = make([]*Job, 0)
)

type Job struct {
	Name   string
	Status bool
	Cron   string
	Func   func()
}

var cronScheduler *cron.Cron

func StartCronJob() {
	addJob()
	// 创建一个支持秒级的 cron 调度器
	cronScheduler = cron.New(cron.WithSeconds())

	for _, job := range JobList {
		if job.Status {
			logger.Logger.Info("添加定时任务" + job.Name + " " + job.Cron)
			cronScheduler.AddFunc(job.Cron, job.Func)
		}
	}

	// 启动 cron 调度器
	cronScheduler.Start()
}

// StopCronJob 停止定时任务
func StopCronJob() {
	if cronScheduler != nil {
		// 停止 cron 调度器
		stop := cronScheduler.Stop()
		// 等待所有任务完成
		<-stop.Done()
	}
}
