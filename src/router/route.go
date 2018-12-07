package router

import (
	"demo/src/handler"
	"demo/src/middleware/jwtpath"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Init(e *gin.Engine) {
	loginController := handler.NewLoginController()
	e.POST("/api/user/login", loginController.Login)
	e.POST("/api/user/logout", loginController.LoginOut)

	api := e.Group("/api", jwtpath.JWTAuth())

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
		userGroup.GET("/info", loginController.UserInfo)

	}

	/**
	task router
	*/
	taskController := handler.NewTaskController()
	taskGroup := api.Group("/task")
	{
		taskGroup.POST("/create", taskController.AddTasks)
	}

	serviceController := handler.NewServiceController()
	serviceGroup := api.Group("/service")
	{
		serviceGroup.GET("/list", serviceController.ServiceList)
	}
	//swagger config
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
