package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type MonitorConf struct {
	Enable         bool   `yaml:"enable"`         // 是否启用上报监控信息
	ReportUrl      string `yaml:"reportUrl"`      // 上报地址
	ReportInterval int    `yaml:"reportInterval"` // 上报时间间隔s
	HttpTimeout    int    `yaml:"httpTimeout"`    // 上报请求超时时间间隔s
}

var Conf MonitorConf

func Init() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	vip := viper.New()
	vip.AddConfigPath(path)         //设置读取的文件路径
	vip.AddConfigPath("/apps/conf") //设置读取的文件路径
	vip.SetConfigName("config")     //设置读取的文件名
	vip.SetConfigType("yaml")       //设置文件的类型
	//尝试进行配置读取
	if err := vip.ReadInConfig(); err != nil {
		panic(err)
	}
	err = vip.Unmarshal(&Conf)
	if err != nil {
		panic(err)
	}
	fmt.Println(Conf)
}
