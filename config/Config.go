package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`

	Database struct {
		Db         string `yaml:""`
		DbHost     string `yaml:"DbHost"`
		DbPort     string `yaml:"DbPort"`
		DbUser     string `yaml:"DbUser"`
		DbPassword string `yaml:"DbPassword"`
		DbName     string `yaml:"DbName"`
	} `yaml:"database"`
}

var Configs Config

func init() {
	data, err := os.ReadFile("Config/config.yml")
	if err != nil {
		panic(fmt.Sprintf("配置文件读取错误,请检查文件路径--%s", err))
	}
	err = yaml.Unmarshal(data, &Configs)
	if err != nil {
		fmt.Println("Unmarshal error:", err)
		panic(fmt.Sprintf("yaml文件提取错误,请检查配置文件--%s", err))
	}
}
