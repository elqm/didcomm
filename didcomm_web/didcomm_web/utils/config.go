package utils

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Port              string `json:"port"`
	BackServerURL     string `json:"bsUrl"`
	WebSocketEndPoint string `json:"wsEndPoint"`
	PathBundle        string `json:"pathBundle"`
	PathHtml          string `json:"pathHtml"`
}

func InitConfig() *Config {
	config := &Config{}
	config.LoadConfig()
	log.Println("설정 파일 정보 불러오기 완료.")
	return config
}

/* 설정 파일(config.json) 정보 불러오기 */
func (config *Config) LoadConfig() {
	file := "config.json"

	path, _ := getFilePath()
	if path == "" {
		return
	}

	log.Println("config 파일 불러오는 중...")

	data, err := os.ReadFile(path + "/config/" + file)
	if err != nil {
		log.Printf("readFile err: %v\n", err)
	}

	err = json.Unmarshal(data, config)
	if err != nil {
		log.Printf("unmarshal err: %v\n", err)
	}

}

/* 현재 작업 디렉토리 경로 반환 (.../DIDCOMM WEB SERVER) */
func getFilePath() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		log.Printf("getPath err: %v\n", err)
	}
	return path, nil
}
