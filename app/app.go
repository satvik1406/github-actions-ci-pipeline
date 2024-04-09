package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"helloworld/domain"
)

var (
	router = gin.Default()
)

func StartingFunction(){
	domain.InitialiseMongoDB(false)
	fmt.Println("DATABASE STARTED")
	routes()
	router.Run(":3000")
}
