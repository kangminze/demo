package router

import (
	"demo/src/handler"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Init(e *gin.Engine) {

	api := e.Group("/api")

	/**
	user router
	*/
	userController := handler.NewUserController()
	userGroup := api.Group("/user")
	{
		userGroup.POST("/create", userController.AddUser)
		userGroup.DELETE("/:id/delete", userController.DeleteUser)
		userGroup.POST("/update/:id", userController.Update)
		userGroup.GET("/details/:id", userController.Details)
		userGroup.GET("/list", userController.ListByPage)

		loginController := handler.NewLoginController()
		userGroup.POST("/login", loginController.Login)
		userGroup.GET("/info", loginController.UserInfo)
		userGroup.POST("/logout", loginController.LoginOut)
	}

	/**
	task router
	*/
	taskController := handler.NewTaskController()
	taskGroup := api.Group("/task")
	{
		taskGroup.POST("/create", taskController.AddTasks)
	}

	//swagger config
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
