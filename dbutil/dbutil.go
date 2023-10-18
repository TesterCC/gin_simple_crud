// dbutil.go
package dbutil

import (
	"context"
	"go_simple_crud/config"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	database   *mongo.Database
	DBEngine   *mongo.Database
)

func init(){
	clientOptions := options.Client().ApplyURI(config.MongoDB)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	DBEngine = client.Database(config.MongoDBName)

}

func ConnectDB(connectionString, dbName, username, password string) (*mongo.Database, error){
	clientOptions := options.Client().ApplyURI(connectionString)
	clientOptions.Auth = &options.Credential{
		Username: username,
		Password: password,
		AuthSource: "admin", // 指定认证数据库
	}

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}


	database = client.Database(dbName)
	return database, nil
}

// 连接到不同的集合
func ConnectToCollection(collectionName string) *mongo.Collection {
	return database.Collection(collectionName)
}

func InsertOne(col *mongo.Collection, document interface{}) (*mongo.InsertOneResult, error) {
	result, err := col.InsertOne(context.Background(), document)
	return result, err
}

// FindOne 函数从指定集合中查找文档
func FindOne(col *mongo.Collection, filter interface{}, result interface{}) error {
	err := col.FindOne(context.Background(), filter).Decode(result)
	return err
}

// UpdateOne 函数更新指定集合中的文档
func UpdateOne(col *mongo.Collection, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	result, err := col.UpdateOne(context.Background(), filter, update)
	return result, err
}

// DeleteOne 函数从指定集合中删除文档
func DeleteOne(col *mongo.Collection, filter interface{}) (*mongo.DeleteResult, error) {
	result, err := col.DeleteOne(context.Background(), filter)
	return result, err
}

// 获取满足过滤条件的文档数量
func CountDocuments(col *mongo.Collection, filter interface{}) (int64, error) {
	count, err := col.CountDocuments(context.Background(), filter)
	if err != nil {
		return 0, err
	}
	return count, nil
}