package utils

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)

func Init() {
	// 获取当前执行文件的目录
	exePath, err := os.Executable()
	if err != nil {
		log.Fatalf("获取执行文件路径失败: %v", err)
	}
	exeDir := filepath.Dir(exePath)

	// 拼接.env文件的相对路径
	envPath := filepath.Join(exeDir, "..", ".env")

	// 加载环境变量
	err = godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("加载 .env 文件失败: %v", err)
	}

}
