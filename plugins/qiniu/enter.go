package qiniu

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"gvb_blog/config"
	"gvb_blog/global"
	"time"
)

// GetToken 获取上传token

func GetToken(q config.QiNiu) string {
	accessKey := q.AccessKey
	secretKey := q.SecretKey
	bucket := q.Bucket
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	return upToken
}

// 获取上传的配置
func getCfg(q config.QiNiu) storage.Config {
	cfg := storage.Config{}
	// 空间的机房
	zone, _ := storage.GetRegionByID(storage.RegionID(q.Zone))
	cfg.Zone = &zone
	// 是否使用https的域名
	cfg.UseHTTPS = false
	// 上传是否使用cdn上传加速
	cfg.UseCdnDomains = false
	return cfg
}

// 上传图片
func UploadImage(data []byte, prefix string, fileName string) (filePath string, err error) {
	// 判断是否使用七牛云
	if !global.Config.QiNiu.Enable {
		return "", errors.New("您没有使用七牛")
	}
	q := global.Config.QiNiu
	if q.AccessKey == "" || q.SecretKey == "" {
		return "", errors.New("请配置accessKey及secretKey")
	}
	if float64(len(data))/1024/1024 > q.Size {
		return "", errors.New("文件超过预定大小")
	}
	upToken := GetToken(q)
	cfg := getCfg(q)

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{},
	}
	dataLen := int64(len(data))
	// 获取当前时间
	now := time.Now().Format("20060102150405")
	key := fmt.Sprintf("%s/%s__%s", prefix, now, fileName)

	err = formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%s", q.CDN, ret.Key), nil
}
