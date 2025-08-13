package mysql

import (
	"gin-scaffold/internal/app/model"
	"os"
	"path/filepath"
	"strings"
)

func Migration() {
	// 自动迁移
	//	 读取所有sql文件
	path := "./data/sql"
	// 使用 filepath.Walk 递归遍历所有目录和文件
	err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 只处理以 .sql 结尾的文件
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".sql") {
			// 读取文件内容
			content, err := os.ReadFile(filePath)
			if err != nil {
				return err
			}

			// 执行sql
			db := GetDB()
			if db == nil {
				return nil
			}

			modelTest := model.JobMq{}
			db.First(&modelTest)
			sql := string(content)
			db.Exec(sql)
		}

		return nil
	})

	if err != nil {
		panic(err)
	}
}
