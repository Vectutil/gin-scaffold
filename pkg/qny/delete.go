package qny

import (
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"net/url"
)

// DeleteFromQiNiuFullPaths 删除七牛云文件
func (q *Qny) DeleteFromQiNiuFullPaths(ctx context.Context, fullPath []string) error {

	mac := qbox.NewMac(q.AccessKey, q.SerectKey)
	zone, _ := storage.GetRegion(q.AccessKey, q.Bucket)
	cfg := storage.Config{
		Zone:          zone,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	bucketManager := storage.NewBucketManager(mac, &cfg)

	for _, fileUrl := range fullPath {
		parsedUrl, err := url.Parse(fileUrl)
		if err != nil {
			fmt.Println("URL 解析失败:", err)
			continue
		}
		// 提取路径部分
		path := parsedUrl.Path
		path = path[1:] // 移除开头的 "/"
		// 截取url中的路径部分
		err = bucketManager.Delete(q.Bucket, path)
		fmt.Println("delete file: ", path)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	return nil
}

func (q *Qny) DeleteFromQiNiuFullPath(ctx context.Context, fullPath string) error {

	mac := qbox.NewMac(q.AccessKey, q.SerectKey)
	zone, _ := storage.GetRegion(q.AccessKey, q.Bucket)
	cfg := storage.Config{
		Zone:          zone,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	bucketManager := storage.NewBucketManager(mac, &cfg)

	parsedUrl, err := url.Parse(fullPath)
	if err != nil {
		fmt.Println("URL 解析失败:", err)
		return err
	}
	// 提取路径部分
	path := parsedUrl.Path
	path = path[1:] // 移除开头的 "/"
	// 截取url中的路径部分
	err = bucketManager.Delete(q.Bucket, path)
	fmt.Println("delete file: ", path)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
