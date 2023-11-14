// dbutil.go
package dbutil

import (
	"context"
	"errors"
	"go_simple_crud/config"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	once     sync.Once
	client   *mongo.Client
	database *mongo.Database
	DBEngine *mongo.Database
)

var (
	ErrNotFound = errors.New("document not found")
)

func init() {
	once.Do(func() {
		clientOptions := options.Client().ApplyURI(config.MongoDB)

		var err error
		client, err := mongo.Connect(context.Background(), clientOptions)
		if err != nil {
			log.Fatal(err)
		}

		DBEngine = client.Database(config.MongoDBName)
	})

}

func ConnectDB(connectionString, dbName, username, password string) (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(connectionString)
	clientOptions.Auth = &options.Credential{
		Username:   username,
		Password:   password,
		AuthSource: "admin", // 指定认证数据库
	}

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return client.Database(dbName), nil
}

// 连接到不同的集合
func ConnectToCollection(collectionName string) *mongo.Collection {
	return database.Collection(collectionName)
}

func InsertOne(ctx context.Context, col *mongo.Collection, document interface{}) (*mongo.InsertOneResult, error) {
	result, err := col.InsertOne(ctx, document)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func FindOne(ctx context.Context, col *mongo.Collection, filter interface{}, result interface{}) error {
	err := col.FindOne(ctx, filter).Decode(result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return ErrNotFound
		}
		return err
	}
	return nil
}

func UpdateOne(ctx context.Context, col *mongo.Collection, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	result, err := col.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func DeleteOne(ctx context.Context, col *mongo.Collection, filter interface{}) (*mongo.DeleteResult, error) {
	result, err := col.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 获取满足过滤条件的文档数量
func CountDocuments(ctx context.Context, col *mongo.Collection, filter interface{}) (int64, error) {
	count, err := col.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}
	return count, nil
}

/*
# 20231114 improve dbutil.go
使用 sync.Once 实现了懒惰单例模式，确保 client 和 DBEngine 只会初始化一次。这样可以避免多次调用 init() 函数创建多个连接。

在 ConnectDB() 函数中，不再使用全局变量 database，而是直接返回连接的数据库实例。这样可以避免多个连接共享同一个全局数据库实例的问题。

对函数参数进行了调整，使用上下文 ctx 作为第一个参数，并返回错误作为第二个参数，以便调用者可以更好地处理错误和取消操作。

添加了 ErrNotFound 错误变量，用于在 FindOne() 函数中表示文档未找到的情况。
*/
