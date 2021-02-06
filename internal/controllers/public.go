package controllers

import (
	"api_server/api/file_server"
	"api_server/api/wechat_server"
	"api_server/pkg/resp"
	"context"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type PublicController struct {
}

func (s *PublicController) UploadFile(c *gin.Context) {
	fileHead, err := c.FormFile("file")
	if err != nil {
		resp.RespParamErr(c, err.Error())
		return
	}
	file, err := fileHead.Open()
	if err != nil {
		resp.RespGeneralErr(c, err.Error())
		return
	}
	defer file.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		resp.RespGeneralErr(c, err.Error())
		return
	}
	client := file_server.GetClient()
	ret, err := client.UploadLocal(context.Background(), &file_server.UploadLocalReq{FileBytes: fileBytes})
	if err != nil {
		resp.RespGeneralErr(c, err.Error())
		return
	}
	if ret.Res.Code != 200 {
		resp.RespGeneralErr(c, ret.Res.Msg)
		return
	}
	resp.RespOk(c, ret)
	return
}

func (s *PublicController) GetShortUrl(c *gin.Context) {
	lurl := c.Query("lurl")
	if lurl == "" {
		resp.RespParamErr(c)
		return
	}
	client := wechat_server.GetClient()
	ret, err := client.MakeShortUrl(c, &wechat_server.MakeShortUrlReq{
		Account: wechat_server.Account_momo_za_huo_pu,
		Url:     lurl,
	})
	if err != nil {
		resp.RespGeneralErr(c, err.Error())
		return
	}
	if ret.Res.Code != 200 {
		resp.RespGeneralErr(c, ret.Res.Msg)
		return
	}
	resp.RespOk(c, ret)
	return
}

func (s *PublicController) GetJssdk(c *gin.Context) {
	url := c.Query("url")
	if url == "" {
		resp.RespParamErr(c)
		return
	}
	client := wechat_server.GetClient()
	ret, err := client.GetJssdk(c, &wechat_server.GetJssdkReq{
		Account: wechat_server.Account_momo_za_huo_pu,
		Url:     url,
	})
	if err != nil {
		resp.RespGeneralErr(c, err.Error())
		return
	}
	if ret.Res.Code != 200 {
		resp.RespGeneralErr(c, ret.Res.Msg)
		return
	}
	resp.RespOk(c, ret)
	return
}
