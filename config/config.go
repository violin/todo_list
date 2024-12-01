package config

import (
	"fmt"
	"gopkg.in/ini.v1"
	"strings"
	"todo_list/model"
)

var (
	AppMode    string
	HttpPort   string
	RedisAddr  string
	RedisPw    string
	RedisDB    string
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
)

func Init() {
	file, err := ini.Load("./config/conf.ini")
	if err != nil {
		fmt.Printf("Fail to parse 'config/conf.ini': %v", err)
	}
	LoadConfig(file)
	path := strings.Join([]string{DbUser, ":", DbPassword, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8"}, "")
	model.Database(path)
}

func LoadConfig(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisPw = file.Section("redis").Key("RedisPw").String()
	RedisDB = file.Section("redis").Key("RedisDB").String()
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassword = file.Section("mysql").Key("DbPassword").String()
	DbName = file.Section("mysql").Key("DbName").String()
}
