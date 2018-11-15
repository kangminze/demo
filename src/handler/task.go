package handler

import (
	"demo/src/model"
	"demo/src/work"
	"github.com/gin-gonic/gin"
	"github.com/sevenNt/wzap"
)


type TaskController struct{}

func NewTaskController() *TaskController {
	return &TaskController{}
}

func (t *TaskController) AddTasks(ctx *gin.Context)  {

	wzap.Debug("this is add tasks")
	task := model.Task{
		ID: 1,
		Content: "this is text tasks",
	}
	work.AddTaskToWork(task)
}
