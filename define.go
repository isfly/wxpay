package wxpay

type unifiedOrder_Req struct {
	AppId          string `xml:"appid"`
	MchId          string `xml:"mch_id"`
	NonceStr       string `xml:"nonce_str"`
	SignType       string `xml:"sign_type"`
	Sign           string `xml:"sign"`
	DeviceInfo     string `xml:"device_info"`
	Body           string `xml:"body"`
	Detail         string `xml:"detail"`
	Attach         string `xml:"attach"`
	OutTradeNo     string `xml:"out_trade_no"`
	FeeType        string `xml:"fee_type"`
	TotalFee       string `xml:"total_fee"`
	SpBillCreateIp string `xml:"spbill_create_ip"`
	TimeStart      string `xml:"time_start"`
	TimeExpire     string `xml:"time_expire"`
	GoodsTag       string `xml:"goods_tag"`
	NotifyUrl      string `xml:"notify_url"`
	TradeType      string `xml:"trade_type"`
	LimitPay       string `xml:"limit_pay"`
	SceneInfo      string `xml:"scene_info"`
}

type queryOrder_Req struct {
	AppId      string `xml:"appid"`
	MchId      string `xml:"mch_id"`
	NonceStr   string `xml:"nonce_str"`
	OutTradeNo string `xml:"out_trade_no"`
	SignType   string `xml:"sign_type"`
	Sign       string `xml:"sign"`
}

type closeOrder_Req struct {
	AppId      string `xml:"appid"`
	MchId      string `xml:"mch_id"`
	NonceStr   string `xml:"nonce_str"`
	OutTradeNo string `xml:"out_trade_no"`
	SignType   string `xml:"sign_type"`
	Sign       string `xml:"sign"`
}

type refund_Req struct {
	AppId         string `xml:"appid"`
	MchId         string `xml:"mch_id"`
	NonceStr      string `xml:"nonce_str"`
	OutTradeNo    string `xml:"out_trade_no"`
	SignType      string `xml:"sign_type"`
	Sign          string `xml:"sign"`
	OutRefundNo   string `xml:"out_refund_no"`
	TotalFee      string `xml:"total_fee"`
	RefundFee     string `xml:"refund_fee"`
	RefundFeeType string `xml:"refund_fee_type"`
	RefundDesc    string `xml:"refund_desc"`
	RefundAccount string `xml:"refund_account"`
}

type queryRefund_Req struct {
	AppId         string `xml:"appid"`
	MchId         string `xml:"mch_id"`
	NonceStr      string `xml:"nonce_str"`
	SignType      string `xml:"sign_type"`
	Sign          string `xml:"sign"`
	OutTradeNo    string `xml:"out_trade_no"`
	TransactionId string `xml:"transaction_id"`
	OutRefundNo   string `xml:"out_refund_no"`
	RefundId      string `xml:"refund_id"`
	OffSet        string `xml:"offset"`
}

type download_Req struct {
	AppId    string `xml:"appid"`
	MchId    string `xml:"mch_id"`
	NonceStr string `xml:"nonce_str"`
	SignType string `xml:"sign_type"`
	Sign     string `xml:"sign"`
	BillDate string `xml:"bill_date"`
	BillType string `xml:"bill_type"`
	TarType  string `xml:"tar_type"`
}

type SIGNTYPE string

const (
	MD5         SIGNTYPE = "MD5"
	HMAC_SHA256 SIGNTYPE = "HMAC-SHA256"
)

type ORDERIDTYPE string

const (
	TRANSACTIONID ORDERIDTYPE = "TRANSACTIONID"
	OUTTRADENO    ORDERIDTYPE = "OUTTRADENO"
	OUTREFUNDNO   ORDERIDTYPE = "OUTREFUNDNO"
	REFUNDID      ORDERIDTYPE = "REFUNDID"
)

type BILLTYPE string

const (
	ALL             BILLTYPE = "ALL"
	SUCCESS         BILLTYPE = "SUCCESS"
	REFUND          BILLTYPE = "REFUND"
	RECHARGE_REFUND BILLTYPE = "RECHARGE_REFUND"
)
