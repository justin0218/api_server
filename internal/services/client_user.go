package services

import (
	"api_server/api/auth_server"
	"api_server/api/user_server"
	"api_server/api/wechat_server"
	"api_server/internal/models/client_user"
	"context"
)

type ClientUser struct {
}

func (s *ClientUser) Login(ctx context.Context, code string) (ret client_user.LoginRes, err error) {
	wechatServer := wechat_server.GetClient()
	authToken, e := wechatServer.GetAuthAccessToken(ctx, &wechat_server.GetAuthAccessTokenReq{
		Account: wechat_server.Account_momo_za_huo_pu,
		Code:    code,
	})
	if e != nil {
		err = e
		return
	}
	wuserInfo, e := wechatServer.GetUserInfo(ctx, &wechat_server.GetUserInfoReq{
		AuthAccessToken: authToken.AccessToken,
		Openid:          authToken.Openid,
	})
	if e != nil {
		err = e
		return
	}
	authServer := auth_server.GetClient()
	userServer := user_server.GetClient()
	oldUser, e := userServer.ClientGetUserByOpenid(ctx, &user_server.ClientGetUserByOpenidReq{
		Openid: wuserInfo.Openid,
	})
	if e != nil {
		err = e
		return
	}
	if oldUser.Code == 404 { //未注册
		newUser, e := userServer.ClientCreateUser(ctx, &user_server.ClientCreateUserReq{
			Openid:   wuserInfo.Openid,
			Nickname: wuserInfo.Nickname,
			Avatar:   wuserInfo.Headimgurl,
		})
		if e != nil {
			err = e
			return
		}
		ret.Openid = wuserInfo.Openid
		ret.Nickname = wuserInfo.Nickname
		ret.Avatar = wuserInfo.Headimgurl
		ret.Uid = newUser.Uid
		tokenRet, e := authServer.CreateToken(ctx, &auth_server.CreateTokenReq{
			Uid:       newUser.Uid,
			TokenType: auth_server.TokenType_CLIENT,
		})
		if e != nil {
			err = e
			return
		}
		ret.Token = tokenRet.Token
		return
	}
	go userServer.ClientUpdateByUid(ctx, &user_server.ClientUpdateByUidReq{
		Uid:      oldUser.Uid,
		Avatar:   wuserInfo.Headimgurl,
		Nickname: wuserInfo.Nickname,
	})

	ret.Openid = wuserInfo.Openid
	ret.Nickname = wuserInfo.Nickname
	ret.Avatar = wuserInfo.Headimgurl
	ret.Uid = oldUser.Uid
	tokenRet, e := authServer.CreateToken(ctx, &auth_server.CreateTokenReq{
		Uid:       oldUser.Uid,
		TokenType: auth_server.TokenType_CLIENT,
	})
	if e != nil {
		err = e
		return
	}
	ret.Token = tokenRet.Token
	return
}
