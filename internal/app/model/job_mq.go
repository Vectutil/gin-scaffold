package model

import (
	"gorm.io/gorm"
	"time"
)

type JobMq struct {
	Id         int64          `json:"id"         orm:"id"           description:"主键"`
	MainUserId int64          `json:"mainUserId" orm:"main_user_id" description:"租户id"`
	ObjId      int64          `json:"objId"      orm:"obj_id"       description:"要实现的id"`
	TargetId   string         `json:"targetId"   orm:"target_id"    description:"目标id"`
	Type       int            `json:"type"       orm:"type"         description:"1舆论"`
	CreatedAt  time.Time      `json:"createdAt"  orm:"created_at"   description:"创建时间"`
	DeletedAt  gorm.DeletedAt `json:"deletedAt"  orm:"deleted_at"   description:"删除时间"`
}
