package main

import (
	"fmt"
	"github.com/minio/minio-go"
	"log"
)

const (
	endpoint        string = "127.0.01:9000"
	accessKeyID     string = "minioadmin"
	secretAccessKey string = "minioadmin"
	useSSL          bool   = false
)

var (
	client *minio.Client
	err    error
)

func init() {
	client, err = minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln("minio连接错误: ", err)
	}
	log.Printf("%#v\n", client)
}

func main() {
	// 1.创建 bucket
	createBucket(BUCKET_NAME)

	// 2.列出所有 bucket
	listBucket()

	//dir, _ := os.Getwd()
	//fmt.Println("dir-->", dir)

	// 3.上传文件到 bucket
	FileUploader(BUCKET_NAME, "图片.jpeg", "./minio/source/图片.jpeg", CONTEXT_TYPE_TEXT)

	// 4.从 bucket 下载文件
	FileGet(BUCKET_NAME, "图片.jpeg", "./minio/download/图片.jpeg")
}

// createBucket 创建 bucket
func createBucket(bucketName string) {
	err = client.MakeBucket(bucketName, "")
	if err != nil {
		// 检查 bucket 是否已经存在
		log.Println("创建bucket错误: ", err)
		exists, _ := client.BucketExists(bucketName)
		if exists {
			log.Printf("bucket: %s已经存在", bucketName)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}
}

// listBucket 查看 bucket 列表
func listBucket() {
	buckets, _ := client.ListBuckets()
	for _, bucket := range buckets {
		fmt.Println("[listBucket] bucket-->", bucket)
	}
}

// FileUploader 文件上传
func FileUploader(bucketName, objectName, filePath, contextType string) {
	n, err := client.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contextType})
	if err != nil {
		log.Println("上传失败：", err)
	}
	log.Printf("Successfully uploaded %s of n %d\n", objectName, n)
}

// FileGet 下载文件
func FileGet(bucketName, objectName, downloadFilePath string) {
	err = client.FGetObject(bucketName, objectName, downloadFilePath, minio.GetObjectOptions{})
	if err != nil {
		log.Println("下载错误: ", err)
	}
}
