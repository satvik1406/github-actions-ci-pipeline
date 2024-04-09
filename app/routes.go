package app

import (
	"helloworld/controllers"
	"os"
)

func routes() {
	appName := os.Getenv("APPNAME")
	if appName == "" {
		appName = "app"
	}
	router.POST("/"+appName+"/messages", controllers.CreateMessage)

	router.PUT("/"+appName+"/messages/:message_id", controllers.UpdateMessage)

	router.GET("/"+appName+"/message", controllers.GetMessage)

	router.GET("/hello", controllers.Hello)

	router.GET("/"+appName+"/hello", controllers.HelloAppName)

	router.GET("/"+appName+"/exit", controllers.Exit)

	router.GET("/"+appName+"/health", controllers.Health)
}