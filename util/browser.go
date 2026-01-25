package util

import (
	"github.com/gin-gonic/gin"
	useragent "github.com/wenlng/go-user-agent"
)

func GetOs(c *gin.Context) string {
	userAgent := c.Request.Header.Get("User-Agent")
	os := useragent.GetOsName(userAgent)
	return os
}

func GetBrowser(c *gin.Context) string {
	userAgent := c.Request.Header.Get("User-Agent")
	browser := useragent.GetBrowserName(userAgent)
	return browser
}
