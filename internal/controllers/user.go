package controllers

import (
	"api_server/api/user_server"
	"api_server/pkg/resp"
	"context"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (s *UserController) Login(c *gin.Context) {
	code := c.Query("code")
	userServer := user_server.GetClient()
	req := &user_server.ClientUserWechatLoginReq{Code: code}
	ret, err := userServer.ClientUserWechatLogin(context.Background(), req)
	if err != nil {
		resp.RespGeneralErr(c, err.Error())
		return
	}
	resp.RespOk(c, ret)
	return
}
