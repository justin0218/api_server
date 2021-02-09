package controllers

import (
	"api_server/internal/services"
	"api_server/pkg/resp"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	clientUser services.ClientUser
}

func (s *UserController) Login(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		resp.RespParamErr(c)
		return
	}
	ret, err := s.clientUser.Login(c, code)
	if err != nil {
		resp.RespGeneralErr(c, err.Error())
		return
	}
	resp.RespOk(c, ret)
	return
}
