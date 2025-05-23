package dao

import (
	"context"
	"gin-scaffold/internal/app/model"
	"gin-scaffold/internal/app/types"
	"gorm.io/gorm"
)

type (
	JobMQDao struct {
		db *gorm.DB
	}
	IJobMQDao interface {
		Create(ctx context.Context, model *model.JobMq)
		SearchOne(ctx context.Context, req *types.GetJobMqOneReq) (model *model.JobMq, err error)
	}
)

func NewJobMQDao(db *gorm.DB) IJobMQDao {
	return &JobMQDao{
		db: db,
	}
}

func (d *JobMQDao) Create(ctx context.Context, model *model.JobMq) {

}

func (d *JobMQDao) SearchOne(ctx context.Context, req *types.GetJobMqOneReq) (model *model.JobMq, err error) {
	query := d.db
	if req.Type != 0 {
		query = query.Where("type = ?", req.Type)
	}
	if req.MainUserId != 0 {
		query = query.Where("main_user_id = ?", req.MainUserId)
	}
	err = query.First(&model).Error
	return
}
