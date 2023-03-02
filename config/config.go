package config

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

const (
	DebugMod     = true
	AppName      = "Zil-Mining-Proxy-Go"
	AppNameLower = "zil-mining-proxy-go"
	AppVersion   = "0.0.1"
)

// Config 代表代理的配置信息
type Config struct {
	LogFolder   string `json:"log_folder"`
	ListenPort  int    `json:"listen_port"`
	APIPort     int    `json:"api_port"`
	StratumPort int    `json:"stratum_port"`
}

func GetCurrentOSLogPath() string {
	defaultLogPath := "./logs"

	if runtime.GOOS == "linux" {
		defaultLogPath = "/var/log"
	}
	return defaultLogPath
}

func GetCurrentOSConfigPath() string {
	defaultLogPath := "./"

	if runtime.GOOS == "linux" {
		defaultLogPath = fmt.Sprintf("/etc/%s", AppNameLower)
	}
	return defaultLogPath
}

// ReadConfig 读取并解析配置文件
func ReadConfig() *Config {
	// 如果配置文件不存在，则生成默认配置文件
	filePath := fmt.Sprintf("%s/%s", GetCurrentOSConfigPath(), "config.json")
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Printf("配置文件 %s 不存在，将使用默认配置生成新的配置文件\n", filePath)

		defaultConfig := Config{
			LogFolder:   GetCurrentOSLogPath(),
			ListenPort:  8080,
			APIPort:     8081,
			StratumPort: 3333,
		}

		bytes, err_json := json.MarshalIndent(defaultConfig, "", "  ")
		if err_json != nil {
			log.Fatalf("无法生成默认配置文件：%s", err)
		}

		// 写入新的配置文件
		filepath.Abs(filepath.Dir(filePath))
		err = os.WriteFile(filePath, bytes, 0644)
		if err != nil {
			log.Fatalf("无法写入新的配置文件：%s", err)
		}
		log.Printf("新的配置文件已生成：%s\n", filePath)
	}

	// 读取配置文件
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("无法打开配置文件：%s", err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("无法读取配置文件：%s", err)
	}

	var config Config
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		log.Fatalf("无法解析配置文件：%s", err)
	}

	return &config
}
