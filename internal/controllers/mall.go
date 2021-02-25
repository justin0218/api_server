package controllers

import (
	"api_server/api/mall_server"
	"api_server/api/user_server"
	"api_server/api/wechat_server"
	"api_server/pkg/resp"
	"fmt"
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
	orderInfo, err := mallServer.CreateOrder(c, req)
	if err != nil {
		resp.RespInternalErr(c, "库存不足")
		return
	}
	userServer := user_server.GetClient()
	userInfo, err := userServer.ClientGetUserByUid(c, &user_server.ClientGetUserByUidReq{
		Uid: req.Uid,
	})
	if err != nil {
		resp.RespInternalErr(c, "不存在的用户")
		return
	}
	wechatServer := wechat_server.GetClient()
	payInfo, err := wechatServer.DoPay(c, &wechat_server.DoPayReq{
		Openid:    userInfo.Openid,
		OrderCode: orderInfo.OrderCode,
		Body:      fmt.Sprintf("%s-%s", orderInfo.GoodsName, orderInfo.SkuName),
		Price:     orderInfo.Price,
		ClientIp:  c.ClientIP(),
		NotifyUrl: "",
		TradeType: "JSAPI",
	})
	if err != nil {
		resp.RespInternalErr(c, err.Error())
		return
	}
	ret := make(map[string]interface{})
	ret["pay_info"] = payInfo
	ret["order_info"] = orderInfo
	resp.RespOk(c, ret)
	return
}
