package alipay

type CreateReq struct {
	Subject     string `json:"subject"`
	TotalAmount string `json:"total_amount"`
}

type PayReq struct {
	OutTradeNo string `json:"out_trade_no"`
}
