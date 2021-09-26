package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Mysql struct {
		Host         string `yml:"host"`
		Port         string `yml:"port"`
		User         string `yml:"user"`
		Pass         string `yml:"pass"`
		DatabaseName string `yml:"databaseName"`
		TableName    string `yml:"tableName"`
		SqlPath      string `yml:"sqlPath"`
	} `yml:"mysql"`
	Email struct {
		User    string `yml:"user"`
		Pass    string `yml:"pass"`
		Host    string `yml:"host"`
		Port    string `yml:"port"`
		Toemail string `yml:"toemail"`
	} `yml:"email"`
}

var GlobalConf Config

func ReadConfig() *Config {
	Conf := new(Config)
	file, err := ioutil.ReadFile("./config/info.yml")
	if err != nil {
		log.Fatalln("读取文件config.yml发生错误", err)
		return nil
	}
	if yaml.Unmarshal(file, Conf) != nil {
		log.Fatalln("解析文件config.yml发生错误")
		return nil
	}
	return Conf
}

func init() {
	GlobalConf = *(ReadConfig())
}
