package component

import (
	"context"
	cli "github.com/milvus-io/milvus-sdk-go/v2/client"
	"log"
	"os"
)

// 访问Milvus数据库的client

var (
	MilvusClient cli.Client
)

func GetMilvusClient(ctx context.Context) *cli.Client {
	var err error
	once.Do(func() {
		MilvusClient, err = cli.NewClient(ctx, cli.Config{
			DBName:  os.Getenv("DB_NAME"),
			Address: os.Getenv("MILVUS_ADDRESS"),
		})
	})

	if err != nil {
		log.Fatalf("init client err:%v", err)
		return nil
	}
	return &MilvusClient
}
