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

func (s *MallController) CreateOrder(c *gin.Context) {
	req := new(mall_server.CreateOrderReq)
	err := c.BindJSON(req)
	if err != nil {
		resp.RespParamErr(c, err.Error())
		return
	}
	req.Uid = getUid(c)
	if req.GoodsId <= 0 || req.SkuId <= 0 || req.BuyNum <= 0 || req.Uid <= 0 || req.Phone == "" || req.Name == "" || req.Province == "" || req.City == "" || req.Region == "" || req.Addr == "" {
		resp.RespParamErr(c)
		return
	}
	mallServer := mall_server.GetClient()
	ret, err := mallServer.CreateOrder(c, req)
	if err != nil {
		resp.RespInternalErr(c, err.Error())
		return
	}
	resp.RespOk(c, ret)
	return
}
