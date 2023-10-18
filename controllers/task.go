package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_simple_crud/models"
	"net/http"
)

type TaskController struct {}

func (t TaskController) Add(c gin.Context) {
	var task models.Task  // todo
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	taskName := task.Name
	taskType := task.Type

	//fmt.Println("[D]", taskName)
	//fmt.Println("[D]", taskType)

	fmt.Printf("[D] %s [D] %d\n", taskName, taskType)

	//if username == "" || password == "" {
	//	ReturnError(c, 40001, "请输入正确的信息")
	//	return
	//}
}