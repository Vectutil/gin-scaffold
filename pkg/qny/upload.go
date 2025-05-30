package qny

import (
	"bytes"
	"context"
	"fmt"
	"gin-scaffold/internal/config"
	"gin-scaffold/pkg/http_call"
	"gin-scaffold/pkg/logger"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Qny struct {
	AccessKey string
	SerectKey string
	Bucket    string
	ImgUrl    string
}

func NewQny() *Qny {
	conf := config.Cfg.Qny
	return &Qny{
		AccessKey: conf.AccessKey,
		SerectKey: conf.SecretKey,
		Bucket:    conf.Bucket,
		ImgUrl:    conf.QnyServer,
	}
}

// UploadToQiNiuWithKey 封装上传图片到七牛云然后返回状态和图片的url
func (q *Qny) UploadToQiNiuWithKey(ctx context.Context, file multipart.File, fileSize int64) (int, string) {
	putPlicy := storage.PutPolicy{
		Scope: q.Bucket,
	}
	mac := qbox.NewMac(q.AccessKey, q.SerectKey)
	upToken := putPlicy.UploadToken(mac)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuadong,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	putExtra := storage.PutExtra{}
	formUploader := storage.NewFormUploader(&cfg)

	ret := storage.PutRet{
		//Key: "goland/test",
	}
	err := formUploader.PutWithoutKey(ctx, &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
		code := 400
		return code, err.Error()
	}
	url := q.ImgUrl + ret.Key
	return 200, url
}

func (q *Qny) UploadToQiNiu(ctx context.Context, path string, fileName string) (int, string) {
	putPolicy := storage.PutPolicy{
		Scope: q.Bucket,
	}
	mac := qbox.NewMac(q.AccessKey, q.SerectKey)
	upToken := putPolicy.UploadToken(mac)
	zone, _ := storage.GetRegion(q.AccessKey, q.Bucket)
	cfg := storage.Config{
		Zone:          zone,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	putExtra := storage.PutExtra{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	key := path
	url := q.ImgUrl + "/" + key
	err := formUploader.PutFile(ctx, &ret, upToken, key, fileName, &putExtra)
	if err != nil {
		if err.Error() == "file exists" {
			return 200, url
		}
		code := 400
		return code, err.Error()
	}
	return 200, url
}

// UploadToQiNiuByUrl 通过URL上传文件到七牛云
func (q *Qny) UploadToQiNiuByUrl(ctx context.Context, path, oldUrl, ext string) (int, string) {
	body, err := http_call.HttpGet(oldUrl, nil)
	if err != nil {
		logger.ErrorLogger.Error("通过URL上传文件到七牛云 失败")
		return 400, ""
	}
	reader := bytes.NewReader([]byte(body))
	key := path + MustEncryptString(oldUrl) + ext
	return q.UploadBinaryToQiNiu(ctx, reader, key)
}

// UploadBinaryToQiNiu 上传二进制数据到七牛云
func (q *Qny) UploadBinaryToQiNiu(ctx context.Context, reader io.Reader, key string) (int, string) {
	if reader == nil {
		return 400, ""
	}

	putPolicy := storage.PutPolicy{
		Scope: fmt.Sprintf("%s:%s", q.Bucket, key),
	}
	mac := qbox.NewMac(q.AccessKey, q.SerectKey)
	upToken := putPolicy.UploadToken(mac)
	zone, _ := storage.GetRegion(q.AccessKey, q.Bucket)
	cfg := storage.Config{
		Zone:          zone,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	putExtra := storage.PutExtra{}

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	url := q.ImgUrl + "/" + key
	err := formUploader.Put(ctx, &ret, upToken, key, reader, -1, &putExtra)
	if err != nil {
		if err.Error() == "file exists" {
			return 200, url
		}
		code := 400
		return code, ""
	}
	return 200, url
}

// GetUploadToken 获取七牛云上传凭证
func (q *Qny) GetUploadToken(ctx context.Context) (token string, imgUrl string) {
	putPolicy := storage.PutPolicy{
		Scope:   q.Bucket,
		Expires: 3600,
	}
	mac := qbox.NewMac(q.AccessKey, q.SerectKey)
	upToken := putPolicy.UploadToken(mac)
	return upToken, q.ImgUrl
}

// ExcelUploadToQiNiu 上传Excel文件到七牛云存储
func (q *Qny) ExcelUploadToQiNiu(ctx context.Context, path string, pathAndFile string, fileName string) (int, string) {

	putPolicy := storage.PutPolicy{
		DeleteAfterDays: 7, //有效期七天
		FileType:        1, //低频存储
		Scope:           q.Bucket,
	}

	mac := qbox.NewMac(q.AccessKey, q.SerectKey)
	upToken := putPolicy.UploadToken(mac)
	zone, _ := storage.GetRegion(q.AccessKey, q.Bucket)
	cfg := storage.Config{
		Zone:          zone,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	putExtra := storage.PutExtra{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	key := path + fileName
	err := formUploader.PutFile(ctx, &ret, upToken, key, pathAndFile, &putExtra)
	if err != nil {
		code := 400
		return code, err.Error()
	}
	url := q.ImgUrl + "/" + key
	return 200, url
}

// DownloadInChunks 分片下载文件并确定文件类型
func (q *Qny) DownloadInChunks(url, fileName, filePath string, chunkSize int) (string, string, error) {
	// 发送 HEAD 请求获取文件大小和内容类型
	resp, err := http.Head(url)
	if err != nil {
		return "", "", fmt.Errorf("发送 HEAD 请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 获取文件大小
	contentLength := resp.ContentLength
	if contentLength == -1 {
		return "", "", fmt.Errorf("无法获取文件大小")
	}

	// 获取内容类型
	contentType := resp.Header.Get("Content-Type")
	if contentType == "" {
		return "", "", fmt.Errorf("无法获取内容类型")
	}

	// 打印内容类型以调试
	fmt.Printf("Content-Type: %s\n", contentType)

	// 从 Content-Disposition 头中提取文件名
	disposition := resp.Header.Get("Content-Disposition")
	var ext string
	if disposition != "" {
		_, params, err := mime.ParseMediaType(disposition)
		if err == nil {
			if filename, ok := params["filename"]; ok {
				ext = filepath.Ext(filename)
				fileName = strings.TrimSuffix(filename, ext)
			}
		}
	}

	// 如果没有从 Content-Disposition 中获取到文件名，则使用 Content-Type 确定文件扩展名
	if ext == "" {
		ext = q.determineFileExtension(contentType)
		if ext == "" {
			ext = ".bin" // 默认使用 .bin
		}
	}

	// 完整的文件路径
	fullFilePath := filepath.Join(filePath, fileName+ext)
	fullFileName := fileName + ext

	// 打开文件以写入
	file, err := os.Create(fullFilePath)
	if err != nil {
		return "", "", fmt.Errorf("创建文件失败: %v", err)
	}
	defer file.Close()

	// 计算分片数量
	numChunks := (contentLength + int64(chunkSize) - 1) / int64(chunkSize)

	for i := int64(0); i < numChunks; i++ {
		start := i * int64(chunkSize)
		end := start + int64(chunkSize) - 1
		if end >= contentLength {
			end = contentLength - 1
		}

		// 发送 GET 请求下载分片
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return "", "", fmt.Errorf("创建请求失败: %v", err)
		}
		req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", start, end))

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return "", "", fmt.Errorf("发送请求失败: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusPartialContent {
			return "", "", fmt.Errorf("请求返回状态码: %d", resp.StatusCode)
		}

		// 读取分片数据
		_, err = io.Copy(file, resp.Body)
		if err != nil {
			return "", "", fmt.Errorf("读取分片数据失败: %v", err)
		}

		fmt.Printf("下载分片 %d/%d 完成\n", i+1, numChunks)
	}

	return fullFilePath, fullFileName, nil
}

// GetQnyUrlIgnoreErr 获取七牛云URL并忽略错误 - 如果获取失败，返回nil和空字符串
func (q *Qny) GetQnyUrlIgnoreErr(ctx context.Context, strUrl, filePath string) (*bytes.Reader, string) {
	resp, err := http_call.HttpGet(strUrl, nil)
	if err != nil {
		return nil, ""
	}

	reader := bytes.NewReader([]byte(resp))
	_, url := q.UploadBinaryToQiNiu(ctx, reader, filePath)
	return reader, url
}

// determineFileExtension 根据内容类型确定文件扩展名
func (q *Qny) determineFileExtension(contentType string) string {
	switch contentType {
	case "image/jpeg", "image/jpg":
		return ".jpg"
	case "image/png":
		return ".png"
	case "image/gif":
		return ".gif"
	case "video/mp4":
		return ".mp4"
	case "video/quicktime":
		return ".mov"
	case "audio/mpeg":
		return ".mp3"
	case "audio/wav":
		return ".wav"
	case "application/pdf":
		return ".pdf"
	case "application/msword":
		return ".doc"
	case "application/vnd.openxmlformats-officedocument.wordprocessingml.document":
		return ".docx"
	case "application/vnd.ms-excel":
		return ".xls"
	case "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet":
		return ".xlsx"
	case "application/octet-stream":
		// 如果无法确定具体类型，可以使用 .bin 或其他通用扩展名
		return ".bin"
	default:
		// 尝试从内容类型中提取扩展名
		parts := strings.Split(contentType, "/")
		if len(parts) == 2 {
			return "." + parts[1]
		}
		return ""
	}
}
