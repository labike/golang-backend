package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type config struct {
	Server    server    `yaml:"server"`
	Dbconfig  dbConfig  `yaml:"db"`
	Redis     redis     `yaml:"redis"`
	Imgupload imgUpload `yaml:"imgUpload"`
	Log       log       `yaml:"log"`
}

// 默认启动配置
type server struct {
	Address string `yaml:"address"`
	Model   string `yaml:"model"`
}

// 数据库配置
type dbConfig struct {
	Dialects string `yaml:"dialects"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
	Charset  string `yaml:"charset"`
	MaxIdle  int    `yaml:"maxIdle"`
	MaxConn  int    `yaml:"maxConn"`
}

// redis
type redis struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
}

// 图片上传
type imgUpload struct {
	UploadDir string `yaml:"uploadDir"`
	Imghost   string `yaml:"imgHost"`
}

// 日志
type log struct {
	Path  string `yaml:"path"`
	Name  string `yaml:"name"`
	Model string `yaml:"model"`
}

var Config *config

func init() {
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &Config)
	fmt.Printf("loaded: %+v\n", Config.Dbconfig)
	if err != nil {
		panic(err)
	}
}
