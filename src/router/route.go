package router

import (
	"demo/src/handler"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Init(e *gin.Engine)  {

	/**
		user router
	 */
	userController := handler.NewUserController()
	userGroup := e.Group("/user")
	{
		userGroup.POST("/create", userController.AddUser)
		userGroup.DELETE("/:id/delete", userController.DeleteUser)
		userGroup.POST("/update/:id" ,userController.Update)
		userGroup.GET("/details/:id", userController.Details)
		userGroup.GET("/list", userController.ListByPage)
	}


	/**
		task router
	 */
	taskController := handler.NewTaskController()
	taskGroup := e.Group("/task")
	{
		taskGroup.POST("/create", taskController.AddTasks)
	}

	//swagger config
	e.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerFiles.Handler))
}
