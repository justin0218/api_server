package controllers

import "github.com/gin-gonic/gin"

func getUid(c *gin.Context) int64 {
	uid, ok := c.Get("uid")
	if !ok {
		return 0
	}
	if uidInt64, ok := uid.(int64); ok {
		return uidInt64
	}
	return 0
}
