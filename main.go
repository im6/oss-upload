package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/spf13/viper"
)

type Configurations struct {
	APPKEY    string
	APPSECRET string
	BUCKET    string
	REGION    string
}

var configuration Configurations

const iconOssDestination string = "1/icons"
const rawFilePath string = "./raw"
const compressedFilePath string = "./compressed"

func loadConfiguration() {
	viper.SetConfigName("oss-upload.env")
	viper.AddConfigPath("../")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		panic(err)
	}
}

func main() {
	loadConfiguration()

	client, err := oss.New(configuration.REGION, configuration.APPKEY, configuration.APPSECRET)
	if err != nil {
		panic(err)
	}
	lsRes, err := client.ListBuckets()
	if err != nil {
		panic(err)
	}
	for _, bucket := range lsRes.Buckets {
		fmt.Println("Buckets:", bucket.Name)
	}
	bucket, err := client.Bucket(configuration.BUCKET)
	options := []oss.Option{
		oss.ContentType("image/svg+xml"),
		oss.ContentEncoding("gzip"),
		oss.CacheControl("public, max-age=31104000"),
	}

	files, err := ioutil.ReadDir(rawFilePath)
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		var fileName = f.Name()
		var destPath = fmt.Sprintf("%s/%s.gz", compressedFilePath, fileName)

		srcFile, err := os.Open(fmt.Sprintf("%s/%s", rawFilePath, fileName))
		if err != nil {
			panic(err)
		}
		defer srcFile.Close()

		dstFile, err := os.Create(destPath)
		if err != nil {
			panic(err)
		}
		defer dstFile.Close()

		zw := gzip.NewWriter(dstFile)
		zw.Name = fileName
		_, err = io.Copy(zw, srcFile)
		if err != nil {
			panic(err)
		}

		zw.Flush()
		if err := zw.Close(); err != nil {
			panic(err)
		}

		err = bucket.PutObjectFromFile(fmt.Sprintf("%s/%s", iconOssDestination, fileName), destPath, options...)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Sent successfully.")
}
