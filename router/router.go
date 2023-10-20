package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go_simple_crud/controllers"
	"go_simple_crud/dbutil"
	"go_simple_crud/models"
	logger "go_simple_crud/pkg"
	"net/http"
)

func Router() *gin.Engine{
	r := gin.Default()

	// 以中间件的形式在路由中调用logger
	r.Use(gin.LoggerWithConfig(logger.LoggerToFile()))
	r.Use(logger.Recover)

	// 配置 sessions_redis的信息
	//store, _ := sessions_redis.NewStore(10, "tcp", config.RedisAddress, "PenTest123", []byte("secret"))
	//r.Use(sessions.Sessions("mysession", store))


	// test router
	// define router directly example
	r.GET("/test", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello World!")
	})
	r.GET("/test/find", func(c *gin.Context) {

		//db, err := dbutil.ConnectDB("mongodb://192.168.80.129:27017", "go_simple_crud", "root", "toor")
		////db, err := dbutil.ConnectDB("mongodb://192.168.200.10:27017", "go_simple_crud", "root", "toor")
		//
		//if err != nil {
		//	log.Fatal(err)
		//	return
		//}
		// global need set in init()
		collection := dbutil.DBEngine.Collection("user")

		//id := c.Param("id")  // "/find/:id"
		id := c.Query("id") // /find?id=xxx
		objID, err := primitive.ObjectIDFromHex(id)

		//filter := bson.M{"_id": id}
		filter := bson.M{"_id": objID}

		fmt.Println("[D]", id)
		fmt.Println("[D]", objID)
		//fmt.Println("[D]", filter)

		var foundUser models.User
		err = dbutil.FindOne(collection, filter, &foundUser)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		fmt.Println("[D]", foundUser)
		c.JSON(200, foundUser)
	})

	//r.POST("/test/insert_simple", func(c *gin.Context) {
	//
	//	db, err := dbutil.ConnectDB("mongodb://192.168.80.129:27017", "go_simple_crud", "root", "toor")
	//	//db, err := dbutil.ConnectDB("mongodb://192.168.200.10:27017", "go_simple_crud", "root", "toor")
	//
	//	if err != nil {
	//		log.Fatal(err)
	//		return
	//	}
	//
	//	//fmt.Println(time.Now().Unix())  // debug
	//
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




    // user router
	userRouter := r.Group("/user")
	// router use controllers   // 大项目推荐用下面这种方式,以防方法名冲突
	{
		userRouter.POST("/register", controllers.UserController{}.Register)
		userRouter.POST("/login", controllers.UserController{}.Login)

		// 路径传参 name
		//user.GET("/info/:name", controllers.UserController{}.GetUserInfo)

	}

	// task router
	taskRouter := r.Group("/task")
	{
		taskRouter.POST("/add", controllers.TaskController{}.Add)
	}

	return r
}