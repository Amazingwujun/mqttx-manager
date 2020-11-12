package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mqttx-manager/entity"
	"net/http"
)

// 用户登入
func SignIn(ctx *gin.Context) {
	user := new(entity.User)
	ctx.Next()
	if err := ctx.ShouldBindJSON(user); err != nil {
		ctx.JSON(http.StatusOK, &entity.Response{
			Code: 400,
			Data: nil,
			Msg:  "参数异常",
		})
		return
	}

	fmt.Printf("收到登入用户信息：%v\n", user)

	// 返回响应
	ctx.JSON(http.StatusOK, &entity.Response{
		Code: 200,
		Data: nil,
		Msg:  "success",
	})
}
