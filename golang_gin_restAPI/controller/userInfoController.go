package controller

import (
	"log"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/takumi616/golang_gin_restAPI/model"
)

func GetAllUserInfo(ctx *gin.Context) {
	userInfoList := model.GetAllUserInfo()

	ctx.Header("Method", ctx.Request.Method)
	statusCode := ctx.Writer.Status()
	ctx.IndentedJSON(statusCode, userInfoList)
}

func GetSingleUserInfo(ctx *gin.Context) {
	param := ctx.Param("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		log.Fatalf("failed to convert String into Int: %v", err)
	}

	userInfo := model.GetSingleUserInfo(id)

	ctx.Header("Method", ctx.Request.Method)
	statusCode := ctx.Writer.Status()
	ctx.IndentedJSON(statusCode, userInfo)
}

func CreateUserInfo(ctx *gin.Context) {
	var userInfo model.UserInfo
	//BindJSON binds the received JSON(ctx.Request *http.Request) to struct of golang.
	ctx.BindJSON(&userInfo)

	returnedId := model.CreateUserInfo(userInfo)

	responseMsg := "Completed creating userInfo record and its id is %d"
	ctx.Header("Method", ctx.Request.Method)
	statusCode := ctx.Writer.Status()
	ctx.String(statusCode, responseMsg, returnedId)
}

func UpdateUserInfo(ctx *gin.Context) {
	param := ctx.Param("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		log.Fatalf("failed to convert String into Int: %v", err)
	}

	var userInfo model.UserInfo
	ctx.BindJSON(&userInfo)

	returnedId := model.UpdateUserInfo(id, userInfo)

	responseMsg := "Completed updating userInfo record and its id is %d"
	ctx.Header("Method", ctx.Request.Method)
	statusCode := ctx.Writer.Status()
	ctx.String(statusCode, responseMsg, returnedId)
}

func DeleteUserInfo(ctx *gin.Context) {
	param := ctx.Param("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		log.Fatalf("failed to convert String into Int: %v", err)
	}

	returnedId := model.DeleteUserInfo(id)
	
	responseMsg := "Completed deleting userInfo record and its id is %d"
	ctx.Header("Method", ctx.Request.Method)
	statusCode := ctx.Writer.Status()
	ctx.String(statusCode, responseMsg, returnedId)
}
