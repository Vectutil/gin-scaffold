package job

import (
	"fmt"
	"gin-scaffold/internal/config"
)

func AddAJob() {
	//AJob
	//talentInfo
	name := "AJob"
	JobList = append(JobList, &Job{
		Status: config.Cfg.Job.JobStatus[name],
		Name:   name,
		Cron:   config.Cfg.Job.JobCron[name],
		Func:   addAJob,
	})
}

func addAJob() {
	fmt.Println("addAJob")
}
