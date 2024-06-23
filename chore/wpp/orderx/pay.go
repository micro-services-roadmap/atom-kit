package orderx

// 支付方式：0->未支付；1->支付宝；2->微信; 3->paypal
const (
	PayTypeAli    = 1
	PayTypeWx     = 2
	PayTypePaypal = 3
)

// 支付状态: 0->待付款；1->已付款；2->已关闭；3->退款异常；4->已退款；
const (
	TradeNotPay    = 0
	TradeSuccess   = 1
	TradeClosed    = 2
	TradeRefundErr = 3
	TradeRefundSuc = 4
)
