package utils

import (
	"fmt"
	"github.com/go-ini/ini"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	Zone       int
	AccessKey  string
	SecretKey  string
	Bucket     string
	QiniuSever string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置我呢见读取错误，请检查文件路径：", err)
	}
	LoadServer(file)
	LoadData(file)
	LoadQiNiu(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("89js82js72")
}
func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassword").MustString("root")
	DbName = file.Section("database").Key("DbName").MustString("blog")
}

func LoadQiNiu(file *ini.File) {
	AccessKey = file.Section("qiniu").Key("AccessKey").String()
	SecretKey = file.Section("qiniu").Key("SecretKey").String()
	Bucket = file.Section("qiniu").Key("Bucket").String()
	QiniuSever = file.Section("qiniu").Key("QiniuSever").String()
	Zone = file.Section("qiniu").Key("Zone").MustInt(2)
}
