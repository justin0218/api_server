package client_user

type LoginRes struct {
	Openid   string `json:"openid"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Uid      int64  `json:"uid"`
	Token    string `json:"token"`
}
