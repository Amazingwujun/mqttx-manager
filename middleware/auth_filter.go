package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 认证过滤器
func AuthFilter(ctx *gin.Context) {
	fmt.Printf("before auth")

	ctx.Next()

	fmt.Printf("after auth")
}
