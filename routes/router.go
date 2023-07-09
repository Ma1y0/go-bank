package routes

import (
	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
    r := gin.Default()

    // Routes
    r.GET("/ping", Ping)

    return r
}


