package component

import (
	"context"
	"github.com/cloudwego/eino-ext/components/model/ark"
	"log"
	"os"
	"time"
)

var (
	chatModel   *ark.ChatModel
	visionModel *ark.ChatModel
)

func GetChatModel(ctx context.Context) (*ark.ChatModel, error) {
	var err error
	timeout := 30 * time.Second
	once.Do(func() {
		chatModel, err = ark.NewChatModel(ctx, &ark.ChatModelConfig{
			APIKey:  os.Getenv("ARK_API_KEY_DOUBAO"),
			Model:   os.Getenv("ARK_MODEL_DOUBAO"),
			Timeout: &timeout,
		})
	})
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return chatModel, nil
}

func GetVisionModel(ctx context.Context) (*ark.ChatModel, error) {
	var err error
	timeout := 30 * time.Second
	once.Do(func() {
		visionModel, err = ark.NewChatModel(ctx, &ark.ChatModelConfig{
			APIKey:  os.Getenv("ARK_VLM_API_KEY"),
			Model:   os.Getenv("ARK_VLM_MODEL"),
			Timeout: &timeout,
		})
	})
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return visionModel, nil
}
