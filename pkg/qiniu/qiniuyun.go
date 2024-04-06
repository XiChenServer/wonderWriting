package qiniu

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"log"
	"mime/multipart"
	"os"
)

var (
	AccessKey string
	SecretKey string
	Bucket    string
	ImgUrl    string
)

func init() {
	// 加载 .env 文件中的环境变量
	if err := godotenv.Load("/home/zwm/GolandProjects/wonderWriting/.env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	// 从环境变量中获取值并赋给相应的变量
	AccessKey = os.Getenv("AccessKey")
	SecretKey = os.Getenv("SecretKey")
	Bucket = os.Getenv("Bucket")
	ImgUrl = os.Getenv("ImgUrl")
}

// UploadToQiNiu 封装上传图片到七牛云然后返回状态和图片的url
func UploadToQiNiu(file multipart.File, fileSize int64, folderPath, fileName string) (string, error) {
	putPlicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, SecretKey)
	upToken := putPlicy.UploadToken(mac)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	putExtra := storage.PutExtra{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {

		return "", err
	}
	url := ret.Key
	fmt.Println(url)
	return ret.Key, nil
}
