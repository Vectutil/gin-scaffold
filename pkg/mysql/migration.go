package mysql

import (
	"gin-scaffold/internal/app/model"
	"os"
	"strings"
)

func Migration() {
	// 自动迁移
	//	 读取所有sql文件
	path := "./data/sql"
	files, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		content := []byte{}
		if strings.HasSuffix(file.Name(), ".sql") {
			// 读取文件内容
			content, err = os.ReadFile(path + "/" + file.Name())
			if err != nil {
				continue
			}
			// 执行sql
			db := GetDB()
			if db == nil {
				continue
			}

			modelTest := model.JobMq{}
			db.First(&modelTest)
			sql := string(content)
			db.Exec(sql)
		}
	}
}
