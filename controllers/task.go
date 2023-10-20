package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go_simple_crud/dbutil"
	"go_simple_crud/models"
	"log"
	"net/http"
	"time"
)

type TaskController struct{}

func (t TaskController) Add(c *gin.Context) {
	var task models.Task

	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	taskName := task.Name
	taskType := task.Type
	taskCmd := task.Command

	//fmt.Println("[D]", taskName)
	//fmt.Println("[D]", taskType)

	fmt.Printf("[D] task name: %s , task type: %d , task command: %s \n", taskName, taskType, taskCmd)

	if taskName == "" {
		ReturnError(c, 40001, "task name is empty")
		return
	}

	if taskType < 1 || taskType > 4 {
		ReturnError(c, 40002, "task type invalid")
		return
	}


	filter := bson.M{"name": taskName}

	collection := dbutil.DBEngine.Collection("task")

	count, err := dbutil.CountDocuments(collection, filter)
	if err != nil {
		// 处理错误
		ReturnError(c, http.StatusInternalServerError, err.Error())
		log.Fatal(err)
	}

	if count > 0 {
		ReturnError(c, 10021, "task name is exist")
	} else {
		// add task info
		newTask := &models.Task {
			Name: taskName,
			Command: taskCmd,
			Status: 0,
			Type: taskType,
			CreatedBy: "admin",
			CreatedAt: time.Now().Unix(),
		}

		result, err := dbutil.InsertOne(collection, newTask)

		if err != nil {
			ReturnError(c, http.StatusInternalServerError, err.Error())
			//c.JSON(500, gin.H{"error": err.Error()})
			//return
		}

		ReturnSuccess(c, http.StatusOK, "request success", gin.H{"insertedID": result.InsertedID})

	}
}
