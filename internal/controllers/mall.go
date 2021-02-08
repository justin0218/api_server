package controllers

import (
	"api_server/api/mall_server"
	"api_server/pkg/resp"
	"github.com/gin-gonic/gin"
	"strconv"
)

type MallController struct {
}

func (s *MallController) GetGoodsInfo(c *gin.Context) {
	goodsId := c.Query("goods_id")
	if goodsId == "" {
		resp.RespParamErr(c)
		return
	}
	goodsIdInt64, err := strconv.ParseInt(goodsId, 10, 64)
	if err != nil {
		resp.RespParamErr(c, err.Error())
		return
	}
	mallServer := mall_server.GetClient()
	ret, err := mallServer.GetGoodsDetail(c, &mall_server.GetGoodsDetailReq{
		GoodsId: goodsIdInt64,
	})
	if err != nil {
		resp.RespInternalErr(c, err.Error())
		return
	}
	resp.RespOk(c, ret)
	return
}
