package core

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"gvb_blog/config"
	"gvb_blog/global"
	"os"
)

const ConfigFilePath = "settings.yaml"

// InitConfig 读取yaml配置
func InitConfig() {
	c := &config.Config{}
	yamlConf, err := os.ReadFile(ConfigFilePath)

	if err != nil {
		global.Log.Fatalf("读取文件失败：", err)
		return
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		global.Log.Fatalf("解析 yaml 文件失败：", err)
		return
	}
	fmt.Println("config yaml load Init Success")
	global.Config = c
}
