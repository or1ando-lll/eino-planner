package component

import (
	"context"
	"eino_planner/utils"
	"github.com/cloudwego/eino-ext/components/indexer/milvus"
	"github.com/cloudwego/eino/components/embedding"
	"github.com/milvus-io/milvus-sdk-go/v2/client"
	"log"
	"os"
)

func GetIndexer(ctx context.Context, milvusClient client.Client, embedder embedding.Embedder) (*milvus.Indexer, error) {
	fields := utils.GetFields()
	indexer, err := milvus.NewIndexer(ctx, &milvus.IndexerConfig{
		Client:     milvusClient,
		Collection: os.Getenv("COLLECTION"),
		Fields:     fields,
		Embedding:  embedder,
	})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return indexer, nil
}
