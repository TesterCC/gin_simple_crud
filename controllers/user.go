package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go_simple_crud/dbutil"
	"go_simple_crud/models"
	"log"
	"net/http"
	"time"
)

// 让每一个controller下面能定义同名函数，使用结构体

type UserController struct{}

func (u UserController) Register(c *gin.Context) {

	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// get json info
	//fmt.Println(user.Username)
	//fmt.Println(user.Email)
	//fmt.Println(user.Roles)

	////获取参数信息 get form_data
	//username := c.DefaultPostForm("username", "")
	//password := c.DefaultPostForm("password", "")
	//confirmPassword := c.DefaultPostForm("confirmPassword", "")

	// get info from json
	username := user.Username
	password := user.Password

	if username == "" || password == "" {
		ReturnError(c, 40001, "请输入正确的信息")
		return
	}

	filter := bson.M{"username": username}

	collection := dbutil.DBEngine.Collection("user")

	count, err := dbutil.CountDocuments(collection, filter)
	if err != nil {
		// 处理错误
		ReturnError(c, http.StatusInternalServerError, err.Error())
		log.Fatal(err)
	}

	if count > 0 {
		ReturnError(c, 10021, "username is exist")
	} else {
		newUser := models.User{
			// 初始化文档数据
			Username: user.Username,
			Email:    user.Email,
			Password: EncryptMd5(user.Password), // md5
			Roles:    user.Roles,                // or []string{"user","admin},
			Status:   uint8(1),
			//CreatedAt: time.Now(),
			//UpdatedAt: time.Now(),
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		}

		result, err := dbutil.InsertOne(collection, newUser)

		if err != nil {
			ReturnError(c, http.StatusInternalServerError, err.Error())
			//c.JSON(500, gin.H{"error": err.Error()})
			//return
		}

		ReturnSuccess(c, http.StatusOK, "request success", gin.H{"insertedID": result.InsertedID})
	}

}

func (u UserController) Login(c *gin.Context) {
	// get login info

	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get info from json
	username := user.Username
	password := user.Password

	if username == "" || password == "" {
		ReturnError(c, 40001, "username or password is empty")
		return
	}

	//fmt.Println("[Debug] user: ", user)

	userCol := dbutil.DBEngine.Collection("user")
	filter := bson.M{"username": username}

	queryUser := models.User{}
    // 为了使 FindOne 函数能够正确地填充查询结果，需要传递结构体的指针，即 &result。这样函数将能够修改结构体的内容并将查询结果填充到结构体中。
	err := dbutil.FindOne(userCol, filter, &queryUser)

	if err != nil {
		// 处理错误
		ReturnError(c, http.StatusInternalServerError, err.Error())
		log.Fatal(err)
	}

	if username != queryUser.Username || EncryptMd5(password) != queryUser.Password {
		ReturnError(c, 40003, "username or password is invalid")
		return
	}

	data := models.UserApi{ID: queryUser.ID, Username: queryUser.Username}
	// todo in real online logic, here need to update session or create jwt token ...
	//fmt.Println("[DDD] ", queryUser)
	ReturnSuccess(c, http.StatusOK, "request success", data)
}
