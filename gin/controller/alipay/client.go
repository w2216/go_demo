package alipay

import (
	"encoding/json"
	"gin/config"
	"gin/model"
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/gopay/pkg/util"
	"github.com/go-pay/gopay/pkg/xlog"
	"github.com/skip2/go-qrcode"
	"net/http"
)

func aliInit() *alipay.Client {
	xlog.Info("GoPay Version: ", gopay.Version)
	// 初始化支付宝客户端
	//    appid：应用ID
	//    privateKey：应用私钥，支持PKCS1和PKCS8
	//    isProd：是否是正式环境
	client, err := alipay.NewClient(config.Appid, config.PrivateKey, config.IsProd)
	if err != nil {
		xlog.Error(err)
		return nil
	}

	// 自定义配置http请求接收返回结果body大小，默认 10MB
	// client.SetBodySize() // 没有特殊需求，可忽略此配置

	// 打开Debug开关，输出日志，默认关闭
	client.DebugSwitch = gopay.DebugOn
	// 设置支付宝请求 公共参数
	//    注意：具体设置哪些参数，根据不同的方法而不同，此处列举出所有设置参数
	client.SetLocation(alipay.LocationShanghai). // 设置时区，不设置或出错均为默认服务器时间
							SetCharset(alipay.UTF8).        // 设置字符编码，不设置默认 utf-8
							SetSignType(alipay.RSA2).       // 设置签名类型，不设置默认 RSA2
							SetReturnUrl(config.ReturnUrl). // 设置返回URL
							SetNotifyUrl(config.NotifyUrl)  // 设置异步通知URL
	// SetAppAuthToken("")             // 设置第三方应用授权

	// 自动同步验签（只支持证书模式）
	// 传入 alipayCertPublicKey_RSA2.crt 内容
	client.AutoVerifySign([]byte("-----BEGIN CERTIFICATE-----\nMIIDqDCCApCgAwIBAgIQICMEIaN9vCw+NXQ8qeHh+zANBgkqhkiG9w0BAQsFADCBkTELMAkGA1UE\nBhMCQ04xGzAZBgNVBAoMEkFudCBGaW5hbmNpYWwgdGVzdDElMCMGA1UECwwcQ2VydGlmaWNhdGlv\nbiBBdXRob3JpdHkgdGVzdDE+MDwGA1UEAww1QW50IEZpbmFuY2lhbCBDZXJ0aWZpY2F0aW9uIEF1\ndGhvcml0eSBDbGFzcyAyIFIxIHRlc3QwHhcNMjMwNDIxMTE1NTA5WhcNMjQwNDIwMTE1NTA5WjB6\nMQswCQYDVQQGEwJDTjEVMBMGA1UECgwM5rKZ566x546v5aKDMQ8wDQYDVQQLDAZBbGlwYXkxQzBB\nBgNVBAMMOuaUr+S7mOWunSjkuK3lm70p572R57uc5oqA5pyv5pyJ6ZmQ5YWs5Y+4LTIwODgxMDIx\nNzk5NzI4NDQwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCxC+q6/Xs0JW75SUMNwxz/\n6VqYG3hekfd8mmkRLtGfsdfZ0XehNf9yhBsICurefUGVo7JaYRI9h+JeYq4GTzQylczR6g+zdtsI\nCR43Uu/8Vm0MnWl6hNDXxN8P8Jde6NxyZJFgMgY25mKAl204koX7NfD5kiPRCP4N7KuTwD0nKNdV\nog/1QivSlv34Io81nfmq1Pomchc1o56umIqDIPC2TK0gauB5tNOBCKo9XhXFmMfkyZf1q8qS541a\nuODBvFNSSpQe/A+doal8RducipmWl+SeKKv+LabL4bIZQN+62YgBbRHwLMf3SvM97ztQMHZEXdNz\nZIB0XvHGE1sN9zh3AgMBAAGjEjAQMA4GA1UdDwEB/wQEAwIE8DANBgkqhkiG9w0BAQsFAAOCAQEA\nkyrZOqncVAEj97MPq3Mg132u8qrWCGm9wXAbRxE/ICmDwEGsaaMYja0HhCMfSx0ZKahEdQb0uDaA\n7mR2HzQLv0kVxNvsogCFuWHB8nGZiks1KHE4s5KUltOQClE/IeMC7xzk3R3ckw3D6FwOSXyHR9Zv\nyU595Nq4bb+3oBVSSscLymUovUm6d4ltZiyg2civW4wpb1zzwwj5De4xJXA7MUDchUPWfLHx5cjz\nFao18C286lfoh1S5O0JVO1L09benETmXMZcWu9TKt+KPI9hSdPWogTHiEGGqpXDUPO9C+nF+VM9P\nMtpoHsDLaR5b16IUqLneOd2C0iFk6cZWDXiHDA==\n-----END CERTIFICATE-----\n-----BEGIN CERTIFICATE-----\nMIIDszCCApugAwIBAgIQIBkIGbgVxq210KxLJ+YA/TANBgkqhkiG9w0BAQsFADCBhDELMAkGA1UE\nBhMCQ04xFjAUBgNVBAoMDUFudCBGaW5hbmNpYWwxJTAjBgNVBAsMHENlcnRpZmljYXRpb24gQXV0\naG9yaXR5IHRlc3QxNjA0BgNVBAMMLUFudCBGaW5hbmNpYWwgQ2VydGlmaWNhdGlvbiBBdXRob3Jp\ndHkgUjEgdGVzdDAeFw0xOTA4MTkxMTE2MDBaFw0yNDA4MDExMTE2MDBaMIGRMQswCQYDVQQGEwJD\nTjEbMBkGA1UECgwSQW50IEZpbmFuY2lhbCB0ZXN0MSUwIwYDVQQLDBxDZXJ0aWZpY2F0aW9uIEF1\ndGhvcml0eSB0ZXN0MT4wPAYDVQQDDDVBbnQgRmluYW5jaWFsIENlcnRpZmljYXRpb24gQXV0aG9y\naXR5IENsYXNzIDIgUjEgdGVzdDCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAMh4FKYO\nZyRQHD6eFbPKZeSAnrfjfU7xmS9Yoozuu+iuqZlb6Z0SPLUqqTZAFZejOcmr07ln/pwZxluqplxC\n5+B48End4nclDMlT5HPrDr3W0frs6Xsa2ZNcyil/iKNB5MbGll8LRAxntsKvZZj6vUTMb705gYgm\nVUMILwi/ZxKTQqBtkT/kQQ5y6nOZsj7XI5rYdz6qqOROrpvS/d7iypdHOMIM9Iz9DlL1mrCykbBi\nt25y+gTeXmuisHUwqaRpwtCGK4BayCqxRGbNipe6W73EK9lBrrzNtTr9NaysesT/v+l25JHCL9tG\nwpNr1oWFzk4IHVOg0ORiQ6SUgxZUTYcCAwEAAaMSMBAwDgYDVR0PAQH/BAQDAgTwMA0GCSqGSIb3\nDQEBCwUAA4IBAQBWThEoIaQoBX2YeRY/I8gu6TYnFXtyuCljANnXnM38ft+ikhE5mMNgKmJYLHvT\nyWWWgwHoSAWEuml7EGbE/2AK2h3k0MdfiWLzdmpPCRG/RJHk6UB1pMHPilI+c0MVu16OPpKbg5Vf\nLTv7dsAB40AzKsvyYw88/Ezi1osTXo6QQwda7uefvudirtb8FcQM9R66cJxl3kt1FXbpYwheIm/p\nj1mq64swCoIYu4NrsUYtn6CV542DTQMI5QdXkn+PzUUly8F6kDp+KpMNd0avfWNL5+O++z+F5Szy\n1CPta1D7EQ/eYmMP+mOQ35oifWIoFCpN6qQVBS/Hob1J/UUyg7BW\n-----END CERTIFICATE-----\n"))

	// 公钥证书模式，需要传入证书，以下两种方式二选一
	// 证书路径
	err = client.SetCertSnByPath("static/appPublicCert.crt", "static/alipayRootCert.crt", "static/alipayPublicCert.crt")
	// 证书内容
	// err := client.SetCertSnByContent("appCertPublicKey bytes", "alipayRootCert bytes", "alipayCertPublicKey_RSA2 bytes")
	if err != nil {
		xlog.Debug("SetCertSn:", err)
		return nil
	}
	return client

}

// 创建订单
func CreateController(ctx *gin.Context) {

	var req CreateReq
	// 在这种情况下，将自动选择合适的绑定
	_ = ctx.ShouldBind(&req)
	if req.Subject == "" || req.TotalAmount == "" {
		xlog.Error(req)
		ctx.JSON(http.StatusOK, "不能为空")
		return
	}

	var Orders model.Orders
	Orders.Subject = req.Subject
	Orders.OutTradeNo = util.RandomString(32)
	Orders.TotalAmount = req.TotalAmount

	db := model.Db
	db.Create(&Orders)
	ctx.JSON(http.StatusOK, gin.H{
		"id":           Orders.ID,
		"out_trade_no": Orders.OutTradeNo,
	})
}

// 支付
func PayController(ctx *gin.Context) {
	var req PayReq
	// 在这种情况下，将自动选择合适的绑定
	_ = ctx.ShouldBind(&req)
	if req.OutTradeNo == "" {
		xlog.Error(req)
		ctx.JSON(http.StatusOK, "不能为空")
		return
	}
	var Orders model.Orders
	db := model.Db
	db.Where("out_trade_no = ?", req.OutTradeNo).First(&Orders)
	client := aliInit()
	bm := make(gopay.BodyMap)

	// 自定义公共参数（根据自己需求，需要独立设置的自行设置，不需要单独设置的，共享client的配置）
	//bm.Set("app_id", "appid")
	//bm.Set("app_auth_token", "app_auth_token")
	//bm.Set("auth_token", "auth_token")

	// biz_content
	bm.SetBodyMap("biz_content", func(bz gopay.BodyMap) {
		bz.Set("subject", Orders.Subject)
		bz.Set("out_trade_no", Orders.OutTradeNo)
		bz.Set("total_amount", Orders.TotalAmount)
	})

	// 当面付 统一收单线下交易预创建接口(用户扫商品收款码)
	aliPsp := new(alipay.TradePrecreateResponse)
	err := client.PostAliPayAPISelfV2(ctx, bm, "alipay.trade.precreate", aliPsp)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Info(aliPsp)
	// 二维码
	f, _ := qrcode.Encode(aliPsp.Response.QrCode, qrcode.Highest, 300)
	_, _ = ctx.Writer.Write(f)

	// 手机网站支付 手机网站支付接口2.0(手机网站支付)
	//bz := make(gopay.BodyMap)
	//bz.Set("subject", Orders.Subject)
	//bz.Set("out_trade_no", Orders.OutTradeNo)
	//bz.Set("total_amount", Orders.TotalAmount)
	//payUrl, err := client.TradeWapPay(ctx, bz)
	//xlog.Info(payUrl)
	//return

	// 手机网站支付 手机网站支付接口2.0(手机网站支付) // 需要提供用户的 auth_code 才能支付
	//bz := make(gopay.BodyMap)
	//bz.Set("subject", Orders.Subject)
	//bz.Set("scene", "bar_code")
	//bz.Set("auth_code", "284490828506449611")
	//bz.Set("out_trade_no", Orders.OutTradeNo)
	//bz.Set("total_amount", Orders.TotalAmount)
	//aliRsp := new(alipay.TradePayResponse)
	//aliRsp, err := client.TradePay(ctx, bz)
	//if err != nil {
	//	xlog.Error(err)
	//	return
	//}
	//xlog.Info(aliRsp)
	//return

	// 当面付 统一收单线下交易预创建接口(用户扫商品收款码) 和第一个相同
	//bz := make(gopay.BodyMap)
	//bz.Set("subject", Orders.Subject)
	//bz.Set("out_trade_no", Orders.OutTradeNo)
	//bz.Set("total_amount", Orders.TotalAmount)
	//aliRsp := new(alipay.TradePrecreateResponse)
	//aliRsp, err := client.TradePrecreate(ctx, bz)
	//if err != nil {
	//	xlog.Error(err)
	//	return
	//}
	//xlog.Info(aliRsp)
	//f, _ := qrcode.Encode(aliRsp.Response.QrCode, qrcode.Highest, 300)
	//_, _ = ctx.Writer.Write(f)
	//return

	////统一收单交易退款接口
	//bz := make(gopay.BodyMap)
	//bz.Set("out_trade_no", Orders.OutTradeNo)
	//bz.Set("refund_amount", "50.00")
	//bz.Set("out_request_no", util.RandomString(32)) // 部分退款必填，默认全额退， 和out_trade_no一致
	//aliRsp := new(alipay.TradeRefundResponse)
	//aliRsp, err := client.TradeRefund(ctx, bz)
	//if err != nil {
	//	xlog.Error(err)
	//	return
	//}
	//xlog.Info(aliRsp)
	//return

	// 统一收单交易退款查询
	//bz := make(gopay.BodyMap)
	//bz.Set("out_trade_no", Orders.OutTradeNo)
	//bz.Set("out_request_no", Orders.OutTradeNo)
	//aliRsp := new(alipay.TradeFastpayRefundQueryResponse)
	//aliRsp, err := client.TradeFastPayRefundQuery(ctx, bz)
	//if err != nil {
	//	xlog.Error(err)
	//	return
	//}
	//xlog.Info(aliRsp)
	//return

	//App支付
	//APP支付接口2.0(APP支付)   ？？？？？
	//bz := make(gopay.BodyMap)
	//bz.Set("subject", Orders.Subject)
	//bz.Set("out_trade_no", Orders.OutTradeNo)
	//bz.Set("total_amount", Orders.TotalAmount)
	//payParam, err := client.TradeAppPay(ctx, bz)
	//if err != nil {
	//	xlog.Error(err)
	//	return
	//}
	//xlog.Info(payParam)
	//return

	//统一收单下单并支付页面接口(电脑网站支付)
	//bz := make(gopay.BodyMap)
	//bz.Set("subject", Orders.Subject)
	//bz.Set("out_trade_no", Orders.OutTradeNo)
	//bz.Set("total_amount", Orders.TotalAmount)
	//payUrl, err := client.TradePagePay(ctx, bz)
	//if err != nil {
	//	xlog.Error(err)
	//	return
	//}
	//xlog.Info(payUrl)
	//return

}

func NotifyController(ctx *gin.Context) {
	req := ctx.Request
	if err := req.ParseForm(); err != nil {
		xlog.Error(err)
		ctx.String(200, "success")
	}
	var form map[string][]string = req.Form
	bm := make(gopay.BodyMap, len(form)+1)
	for k, v := range form {
		if len(v) == 1 {
			bm.Set(k, v[0])
		}
	}
	//map[
	//app_id:2016101700707168
	//auth_app_id:2016101700707168
	//buyer_id:2088102180033471
	//buyer_logon_id:tnp***@sandbox.com
	//buyer_pay_amount:0.01
	//charset:utf-8
	//fund_bill_list:[{"amount":"0.01","fundChannel":"ALIPAYACCOUNT"}]
	//gmt_create:2023-04-22 14:58:57
	//gmt_payment:2023-04-22 14:59:11
	//invoice_amount:0.01
	//notify_id:2023042200222145913033470526456433
	//notify_time:2023-04-22 14:59:13
	//notify_type:trade_status_sync
	//out_trade_no:moEkYPKV1hIMES7ctFuEzbwREn2nhdP0
	//point_amount:0.00 receipt_amount:0.01
	//seller_email:mgrgjd4404@sandbox.com
	//seller_id:2088102179972844
	//sign:cQypv0eFK4yKtRbTQW/UwN7Pz0Hc1r80bpomM0chrH1Xn+Nuj9UsQ35uYcpIbv4mc5JzkKT9Hd89j+OomDEDgZYzuuoOdSLHTqvfPiHgsutVfSYwyWGvtw9xErxrlufhLkG3P1bxsYbdFCcuecFwUA3fiTfr9oIqmohdyW/qoemgLWBAdSk+R8TDwoGtyhAGBy1mDn+W8oIi+ZOD7VMvG/SULOqATlPbr0kvQeCEb3RhwDgO+OzeabhMHNvnrqBy7yEKPC+a6vEbk+8E1H0lk7WF+dGkPzJYludrzjldejhhtqI4vvsCU3HYbRrRJwsmPWJ9MJQjzxrOM++PbtFiQg==
	//sign_type:RSA2
	//subject:0.01元的大啤酒
	//total_amount:0.01
	//trade_no:2023042222001433470502666906
	//trade_status:TRADE_SUCCESS
	//version:1.0
	//]
	xlog.Info(bm)

	// 异步验签
	ok, err := alipay.VerifySignWithCert("static/alipayPublicCert.crt", bm)
	if !ok {
		xlog.Error(err)
		ctx.String(200, "success")
		return
	} else {
		xlog.Info("修改订单")
		DoOrders(bm)
		ctx.String(200, "success")
	}

}

// 修改订单
func DoOrders(bm gopay.BodyMap) {

	OutTradeNo := bm.Get("out_trade_no")
	BuyerId := bm.Get("buyer_id")
	BuyerLogonId := bm.Get("buyer_logon_id")
	BuyerPayAmount := bm.Get("buyer_pay_amount")
	GmtCreate := bm.Get("gmt_create")
	GmtPayment := bm.Get("gmt_payment")
	GmtRefund := bm.Get("gmt_refund")
	NotifyId := bm.Get("notify_id")
	NotifyTime := bm.Get("notify_time")
	NotifyType := bm.Get("notify_type")
	TradeNo := bm.Get("trade_no")
	TradeStatus := bm.Get("trade_status")
	jsonBm, _ := json.Marshal(bm)
	NotifyAll := string(jsonBm)
	Status := 0
	if TradeStatus == "TRADE_SUCCESS" {
		Status = 1
	} else {
		Status = 2
	}

	db := model.Db
	db.Model(&model.Orders{}).Where("out_trade_no = ?", OutTradeNo).Updates(model.Orders{
		BuyerId:        BuyerId,
		BuyerLogonId:   BuyerLogonId,
		BuyerPayAmount: BuyerPayAmount,
		GmtCreate:      GmtCreate,
		GmtPayment:     GmtPayment,
		GmtRefund:      GmtRefund,
		NotifyId:       NotifyId,
		NotifyType:     NotifyType,
		NotifyTime:     NotifyTime,
		TradeNo:        TradeNo,
		TradeStatus:    TradeStatus,
		NotifyAll:      NotifyAll,
		Status:         int64(Status),
	})
}
