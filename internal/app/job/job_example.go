package job

import (
	"gin-scaffold/internal/config"
	"gin-scaffold/pkg/logger"
	"time"
)

func AddExampleJob() {
	//AJob
	//talentInfo
	name := "examplejob" // 必须全小写不能采用驼峰命名
	JobList = append(JobList, &Job{
		Status: config.Cfg.Job.JobStatus[name],
		Name:   name,
		Cron:   config.Cfg.Job.JobCron[name],
		Func:   addExampleJob,
	})
}

func addExampleJob() {
	logger.Logger.Info(time.Now().Format("2006-01-02 15:04:05") + " : addExampleJob")
}

//{"level":"INFO","time":"2025-05-24T09:46:20.000+0800","caller":"job/job_a.go:22","msg":"2025-05-24 09:46:20 : addExampleJob"}
//{"level":"INFO","time":"2025-05-24T09:46:30.001+0800","caller":"job/job_a.go:22","msg":"2025-05-24 09:46:30 : addExampleJob"}
