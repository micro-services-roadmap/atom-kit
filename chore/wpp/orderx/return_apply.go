package orderx

const (
	RaUnhandled = 0 // 未处理
	RaReturning = 1 // 退货中
	RaReturned  = 2 // 已收货完成
	RaRejected  = 3 // 已拒绝

	RaRefunding = 4 // 退款ing
	RaRefunded  = 5 // 已退款
	RaRefundErr = 6 // 退款异常
	RaClosed    = 7 // 已取消
)
