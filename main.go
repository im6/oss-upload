package main

import (
	"fmt"

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
}
