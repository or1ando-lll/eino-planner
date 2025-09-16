package component

import (
	"context"
	"github.com/cloudwego/eino-ext/components/embedding/ark"
	"os"
	"sync"
)

var (
	embedder *ark.Embedder
	once     sync.Once // 确保初始化只执行一次
)

func GetEmbedder(ctx context.Context) (*ark.Embedder, error) {
	var err error
	once.Do(func() {
		// 只在首次调用时执行
		embedder, err = ark.NewEmbedder(ctx, &ark.EmbeddingConfig{
			APIKey: os.Getenv("ARK_API_KEY"),
			Model:  os.Getenv("EMBEDDER"),
		})
	})
	if err != nil {
		return nil, err
	}
	return embedder, nil
}
