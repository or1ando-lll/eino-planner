package utils

import (
	"github.com/joho/godotenv"
	"github.com/milvus-io/milvus-sdk-go/v2/entity"
	"log"
	"os"
	"path/filepath"
)

// EnvInit 加载环境变量
func EnvInit() {
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

// GetFields 获取向量数据哭的表结构
func GetFields() []*entity.Field {
	//定义表的字段格式
	fields := []*entity.Field{
		{
			Name:     "id",
			DataType: entity.FieldTypeVarChar,
			TypeParams: map[string]string{
				"max_length": "256",
			},
			PrimaryKey: true,
		},
		{
			Name:     "vector",
			DataType: entity.FieldTypeBinaryVector,
			TypeParams: map[string]string{
				"dim": "81920",
			},
		},
		{
			Name:     "content",
			DataType: entity.FieldTypeVarChar,
			TypeParams: map[string]string{
				"max_length": "8192",
			},
		},
		{
			Name:     "metadata",
			DataType: entity.FieldTypeJSON,
		},
	}
	return fields
}
