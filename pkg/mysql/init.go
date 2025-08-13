package mysql

import (
	"context"
	"fmt"
	"gin-scaffold/internal/config"
	"gin-scaffold/internal/middleware/metadata"
	sys_logger "gin-scaffold/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

var dbMap *gorm.DB

func InitMysql() {
	//
	jtDB := config.Cfg.Mysql

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		jtDB.User, jtDB.Password, jtDB.Host, jtDB.Port, jtDB.Database)

	// 配置日志
	//newLogger := logger.New(
	//	log.New(os.Stdout, "\r\n", log.LstdFlags),
	//	logger.Config{
	//		SlowThreshold:             time.Second, // 慢查询阈值
	//		LogLevel:                  logger.Info, // 开发环境用Info，生产用Error
	//		IgnoreRecordNotFoundError: true,        // 忽略记录不存在错误
	//		Colorful:                  true,        // 启用彩色日志
	//	},
	//)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false, // 保持事务
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，例如 `User` -> 表名为 `user`，而不是默认的 `users`
		},
		FullSaveAssociations:                     false,                                                          // 禁用级联保存
		Logger:                                   sys_logger.NewGormLogger(1 * time.Second).LogMode(logger.Info), // 自定义日志
		NowFunc:                                  nil,                                                            // 使用默认时间函数
		DryRun:                                   false,                                                          // 非调试模式
		PrepareStmt:                              true,                                                           // 启用预编译缓存
		DisableAutomaticPing:                     true,                                                           // 启用自动ping检查连接
		DisableForeignKeyConstraintWhenMigrating: true,                                                           // 迁移时禁用外键约束
		AllowGlobalUpdate:                        false,                                                          // 禁止无WHERE的全局更新
		QueryFields:                              true,                                                           // 显式指定查询字段
		CreateBatchSize:                          1000,                                                           // 批量插入大小
	})
	if err != nil {
		panic("failed to connect database")
	}

	// 检测是否正确连接

	// 配置连接池
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get underlying sql.DB")
	}
	sqlDB.SetMaxOpenConns(200) // 最大打开连接数
	sqlDB.SetMaxIdleConns(25)  // 最大空闲连接数
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
	//db.Logger = logger.Default.LogMode(logger.Silent)
	setupGlobalHooks(db)
	dbMap = db
}

func GetDB() *gorm.DB {
	return dbMap
}

// GetTrans 开启事务并返回事务句柄和提交函数
func GetTrans() (*gorm.DB, func(err error)) {
	tx := dbMap.Begin()
	commit := func(err error) {
		if r := recover(); r != nil {
			tx.Rollback()
			return
		}

		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}
	return tx, commit
}

func setupGlobalHooks(db *gorm.DB) {
	//db.Callback().Create().Before("gorm:before_create").Register("before_create_hook", beforeCreateHook)
	//db.Callback().Create().Before("gorm:before_update").Register("before_update_hook", beforeUpdateHook)
	db.Callback().Delete().Before("gorm:before_delete").Register("before_delete_hook", beforeDeleteHook)
}

//	func beforeCreateHook(db *gorm.DB) {
//		if db.Error != nil {
//			return
//		}
//		if db.Statement.Schema != nil {
//			idField := db.Statement.Schema.PrimaryFields[0]
//			if idField.DataType == "uint32" {
//				_, zero := idField.ValueOf(db.Statement.Context, db.Statement.ReflectValue)
//				if zero {
//					idField.Set(db.Statement.Context, db.Statement.ReflectValue, uuid.New().ID())
//					db.Statement.SetColumn("created_at", time.Now().Format("2006-01-02 15:04:05"))
//					db.Statement.AddClause(clause.OnConflict{
//						DoNothing: true,
//					})
//				}
//			}
//			//db.Statement.SetColumn("id", uuid.New().String())
//			//db.Statement.SetColumn("created_at", time.Now().Format("2006-01-02 15:04:05"))
//		}
//	}
//
//	func beforeUpdateHook(db *gorm.DB) {
//		if db.Error != nil {
//			return
//		}
//		if db.Statement.Schema != nil {
//			db.Statement.SetColumn("updated_at", time.Now().Format("2006-01-02 15:04:05"))
//		}
//	}
func beforeDeleteHook(db *gorm.DB) {
	if db.Error != nil {
		return
	}
	if db.Statement.Schema != nil {
		//db.Statement.SetColumn("deleted_at", time.Now().Format("2006-01-02 15:04:05"))
		if ctx, ok := db.Statement.Context.(context.Context); ok {
			db.Statement.SetColumn("deleted_by", metadata.GetUserId(ctx))
		}
	}
}
