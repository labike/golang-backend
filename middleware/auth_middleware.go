package middleware

import (
	"go-admin/common/constant"
	"go-admin/common/result"
	"go-admin/pkg/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			result.Failed(c, int(result.ApiCode.NOAUTH), result.ApiCode.GetMessage(result.ApiCode.NOAUTH))
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			result.Failed(c, int(result.ApiCode.AUTHFORMATERROR), result.ApiCode.GetMessage(result.ApiCode.AUTHFORMATERROR))
			c.Abort()
			return
		}
		// 校验token
		mc, err := jwt.ValidateToken(parts[1])
		if err != nil {
			result.Failed(c, int(result.ApiCode.TOKENINLAILD), result.ApiCode.GetMessage(result.ApiCode.TOKENINLAILD))
			c.Abort()
			return
		}
		c.Set(constant.ContextKeyUserObj, mc)
		c.Next()
	}
}
