package wx

type WxpayReq struct {
	Appid         string `xml:"appid"`
	BankType      string `xml:"bank_type"`
	CashFee       string `xml:"cash_fee"`
	FeeType       string `xml:"fee_type"`
	IsSubscribe   string `xml:"is_subscribe"`
	MchId         string `xml:"mch_id"`
	NonceStr      string `xml:"nonce_str"`
	Openid        string `xml:"openid"`
	OutTradeNo    string `xml:"out_trade_no"`
	ResultCode    string `xml:"result_code"`
	ReturnCode    string `xml:"return_code"`
	Sign          string `xml:"sign"`
	TimeEnd       string `xml:"time_end"`
	TotalFee      int    `xml:"total_fee"`
	TradeType     string `xml:"trade_type"`
	TransactionId string `xml:"transaction_id"`
}

type WxpayRes struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
}
