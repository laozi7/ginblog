package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode string
	HttpPost string
	JwtKey string

	Db string
	DbHost string
	DbPort string
	DbUser string
	DbPassWord string
	DbName string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("config.ini文件不存在")
	}
	LoadServer(file)
	LoadData(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPost = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("laozi")
}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").String()
	DbHost = file.Section("database").Key("DbHost").String()
	DbPort = file.Section("database").Key("DbPort").String()
	DbUser = file.Section("database").Key("DbUser").String()
	DbPassWord = file.Section("database").Key("DbPassWord").String()
	DbName = file.Section("database").Key("DbName").String()
}


