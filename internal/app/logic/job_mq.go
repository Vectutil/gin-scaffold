package logic

import (
	"context"
	"gin-scaffold/internal/app/dao"
	"gin-scaffold/internal/app/model"
	"gin-scaffold/internal/app/types"
	"gin-scaffold/pkg/mysql"
)

type (
	JobMQLogic struct {
		jobMqDao dao.IJobMQDao
	}
	IJobMQLogic interface {
		SearchOne(ctx context.Context, req *types.GetJobMqOneReq) (model *model.JobMq, err error)
	}
)

func NewJobMQLogic() IJobMQLogic {
	return &JobMQLogic{}
}

func (l *JobMQLogic) Create(ctx context.Context, model *model.JobMq) {

}

func (l *JobMQLogic) SearchOne(ctx context.Context, req *types.GetJobMqOneReq) (model *model.JobMq, err error) {
	l.jobMqDao = dao.NewJobMQDao(mysql.GetDB())
	return l.jobMqDao.SearchOne(ctx, req)
}
