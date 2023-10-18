// main.go
package main

import (
	"go_simple_crud/router"
)

func main() {
	//r := gin.Default()
	//
	//// Connection to MongoDB
	//db, err := dbutil.ConnectDB("mongodb://192.168.80.129:27017", "go_simple_crud", "root", "toor")
	////db, err := dbutil.ConnectDB("mongodb://192.168.200.10:27017", "go_simple_crud", "root", "toor")
	//
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//
	////"mongodb://root:toor@192.168.80.129:27017/admin"
	//
	//// test router
	//r.POST("/test/insert_simple", func(c *gin.Context) {
	//
	//	fmt.Println(time.Now().Unix())
	//	// 获取或创建集合
	//	collection := db.Collection("user")
	//
	//	newUser := models.User{
	//		// 初始化文档数据
	//		Username: "admin",
	//		Email:    "admin@admin.com",
	//		Password: "PenTest123",
	//		Roles:    []string{"admin"}, // or []string{"user","admin},
	//		//CreatedAt: time.Now(),
	//		//UpdatedAt: time.Now(),
	//		CreatedAt: time.Now().Unix(),
	//		UpdatedAt: time.Now().Unix(),
	//	}
	//
	//	result, err := dbutil.InsertOne(collection, newUser)
	//
	//	if err != nil {
	//		c.JSON(500, gin.H{"error": err.Error()})
	//		return
	//	}
	//
	//	c.JSON(200, gin.H{"message": "Document inserted", "insertedID": result.InsertedID})
	//})
	//
	//r.GET("/test/find", func(c *gin.Context) {
	//
	//	collection := db.Collection("user")
	//
	//	//id := c.Param("id")  // "/find/:id"
	//	id := c.Query("id") // /find?id=xxx
	//	objID, err := primitive.ObjectIDFromHex(id)
	//
	//	//filter := bson.M{"_id": id}
	//	filter := bson.M{"_id": objID}
	//
	//	fmt.Println("[D]", id)
	//	fmt.Println("[D]", objID)
	//	fmt.Println("[D]", filter)
	//
	//	var foundUser models.User
	//	err = dbutil.FindOne(collection, filter, &foundUser)
	//	if err != nil {
	//		c.JSON(500, gin.H{"error": err.Error()})
	//		return
	//	}
	//
	//	fmt.Println("[D]", foundUser)
	//	c.JSON(200, foundUser)
	//})


	r := router.Router()
	r.Run(":8888")
}
