package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/minio/minio-go/v6"
)

func main() {
	endpoint := "192.168.0.4:9000"
	accessKeyID := "minioadmin"
	secretAccessKey := "minioadmin"
	useSSL := false

	// 初使化 minio client对象。
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	// 创建一个叫mymusic的存储桶。
	bucketName := "mymusic"
	location := "us-east-1"

	err = minioClient.MakeBucket(bucketName, location)
	if err != nil {
		// 检查存储桶是否已经存在。
		exists, err := minioClient.BucketExists(bucketName)
		if err == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	}
	log.Printf("Successfully created %s\n", bucketName)

	// 上传一个zip文件。
	objectName := "ea.zip"
	filePath := "/Users/yangtao/tmp/ea.zip"
	contentType := "application/zip"

	// 使用FPutObject上传一个zip文件。
	n, err := minioClient.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, n)
	if object, err := minioClient.GetObject(bucketName, objectName, minio.GetObjectOptions{}); err != nil {
		log.Fatalln(err)
	} else {
		localFile, err := os.Create("/Users/yangtao/tmp" + objectName)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer os.Remove("/Users/yangtao/tmp/" + objectName)
		if _, err = io.Copy(localFile, object); err != nil {
			fmt.Println(err)
			return
		}

	}

	log.Println("Successfully saved filename.csv.")
}
