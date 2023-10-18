package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go_simple_crud/dbutil"
	"go_simple_crud/models"
	"log"
	"net/http"
	"time"
)

type UserApi struct {
	ID       primitive.ObjectID `json:"id"`
	Username string             `json:"username"`
}

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
		log.Fatal(err)
		ReturnError(c, http.StatusInternalServerError, err.Error())
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

//user, err := models.GetUserInfoByUsername(username)
//if user.ID != 0 {
//	ReturnError(c, 40003, "用户名已存在")
//	return
//}
//_, err = models.AddUser(username, EncryptMd5(password))
//if err != nil {
//	ReturnError(c, 40004, "注册失败，请重试")
//	return
//}

//ReturnSuccess(c, 0, "success", "")
