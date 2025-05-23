package job

import (
	"context"
	"fmt"
	"gin-scaffold/internal/app/logic"
	"gin-scaffold/internal/app/types"
	"github.com/robfig/cron/v3"
)

var cronScheduler *cron.Cron

func StartCronJob() {
	// 创建一个支持秒级的 cron 调度器
	cronScheduler = cron.New(cron.WithSeconds())

	jobLogic := logic.NewJobMQLogic()

	// 添加每秒执行一次的定时任务
	_, err := cronScheduler.AddFunc("@every 1s", func() {
		model, err := jobLogic.SearchOne(context.Background(), &types.GetJobMqOneReq{
			Type:       1,
			MainUserId: 1000000,
		})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(model)
	})
	if err != nil {
		fmt.Println("添加定时任务失败:", err)
		return
	}

	// 启动 cron 调度器
	cronScheduler.Start()
	fmt.Println("定时任务启动")
}

// StopCronJob 停止定时任务
func StopCronJob() {
	if cronScheduler != nil {
		// 停止 cron 调度器
		stop := cronScheduler.Stop()
		// 等待所有任务完成
		<-stop.Done()
		fmt.Println("定时任务已停止")
	}
}
