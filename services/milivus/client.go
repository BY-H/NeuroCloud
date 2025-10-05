package milvus

import (
	"context"
	"log"
	"sync"

	"github.com/milvus-io/milvus-sdk-go/v2/client"
)

var (
	milvusClient client.Client
	once         sync.Once
)

// Init 初始化 Milvus 客户端
func Init(addr string) client.Client {
	once.Do(func() {
		c, err := client.NewGrpcClient(context.Background(), addr)
		if err != nil {
			log.Fatalf("failed to connect milvus: %v", err)
		}
		milvusClient = c
	})
	return milvusClient
}

// GetClient 获取 Milvus 客户端
func GetClient() client.Client {
	if milvusClient == nil {
		log.Fatal("milvus client is not initialized, call Init() first")
	}
	return milvusClient
}
