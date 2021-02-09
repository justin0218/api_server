package middleware

import (
	"api_server/api/auth_server"
	"api_server/pkg/resp"
	"api_server/store"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		if new(store.Config).Get().Runmode == "debug" {
			c.Next()
			return
		}
		token := c.Query("token")
		uid := c.Query("uid")
		uidInt64, err := strconv.ParseInt(uid, 10, 64)
		if err != nil {
			resp.RespCode(c, http.StatusUnauthorized, "未授权")
			c.Abort()
			return
		}
		conn := auth_server.GetClient()
		req := &auth_server.VerifyTokenReq{}
		req.TokenType = auth_server.TokenType_CLIENT
		req.Uid = uidInt64
		req.Token = token
		ret, err := conn.VerifyToken(context.Background(), req)
		if err != nil {
			resp.RespCode(c, http.StatusUnauthorized, "未授权:"+err.Error())
			c.Abort()
			return
		}
		//if ret.TokenError == auth_server.TokenError_EXPIRED || ret.TokenError == auth_server.TokenError_USER_MATCH {
		//	resp.RespCode(c, http.StatusUnauthorized, ret.TokenError)
		//	c.Abort()
		//	return
		//}
		c.Set("uid", ret.Uid)
		c.Next()
	}
}
