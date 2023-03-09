package main

import (
    "github.com/gin-gonic/gin"
	"github.com/takumi616/golang_gin_restAPI/controller"
)

func main() {
    router := gin.Default()

    router.GET("/api/userInfos", controller.GetAllUserInfo)
    router.GET("/api/userInfos/:id", controller.GetSingleUserInfo)
    router.POST("api/userInfos", controller.CreateUserInfo)
	router.PUT("/api/userInfos/:id", controller.UpdateUserInfo)
    router.DELETE("api/userInfos/:id", controller.DeleteUserInfo)

    router.Run("localhost:8000")
}

