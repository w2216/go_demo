package model

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

// 用户表
type Admin struct {
	ID           uint `gorm:"primaryKey"`
	Name         string
	Email        string
	Password     string
	Age          uint8
	Birthday     string
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// 登录记录表
type LoginLog struct {
	ID        uint `gorm:"primaryKey"`
	Aid       uint
	CreatedAt int
}

type Orders struct {
	ID              int64  `gorm:"primaryKey" json:"id"`
	OutTradeNo      string `json:"out_trade_no"`      // 必选 商户网站唯一订单号
	TotalAmount     string `json:"total_amount"`      // 必选 订单总金额 单位为元 0.01
	Subject         string `json:"subject"`           // 必选 订单标题
	TimeExpire      string `json:"time_expire"`       // 2016-12-31 10:05:00
	PassbackParams  string `json:"passback_params"`   // 公用回传参数
	MerchantOrderNo string `json:"merchant_order_no"` // 商户原始订单号
	Status          int64  `json:"status"`            // 订单状态 0,创建 1,支付成功 2.支付失败

	BuyerId        string `json:"buyer_id"`         // 买家支付宝账号对应的支付宝唯一用户号
	BuyerLogonId   string `json:"buyer_logon_id"`   // 买家支付宝账号
	BuyerPayAmount string `json:"buyer_pay_amount"` // 用户在交易中支付的金额
	GmtCreate      string `json:"gmt_create"`       // 该笔交易创建的时间
	GmtPayment     string `json:"gmt_payment"`      // 该笔交易 的买家付款时间
	GmtRefund      string `json:"gmt_refund"`       // 该笔交易的退款时间
	NotifyId       string `json:"notify_id"`        // 通知校验 ID
	NotifyTime     string `json:"notify_time"`      // 通知的发送时间
	NotifyType     string `json:"notify_type"`      // 通知的类型
	TradeNo        string `json:"trade_no"`         // 支付宝交易凭证号
	// TRADE_FINISHED交易完成true（触发通知）
	// TRADE_SUCCESS支付成功true（触发通知）
	// WAIT_BUYER_PAY交易创建false（不触发通知）
	// TRADE_CLOSED交易关闭true（触发通知）
	TradeStatus string `json:"trade_status"`
	NotifyAll   string `json:"notify_all"`
}

type OrderDetails struct {
	ID             int64  `gorm:"primaryKey" json:"id"`
	OrderId        int64  `json:"order_id"`
	GoodsId        string `json:"goods_id"`        // 商品编号 必填
	GoodsName      string `json:"goods_name"`      // 商品名称 必填
	Quantity       int64  `json:"quantity"`        // 商品数量 必填
	Price          string `json:"price"`           // 商品单价 必填
	GoodsCategory  string `json:"goods_category"`  // 商品类目
	CategoriesTree string `json:"categories_tree"` // 商品类目树
	Body           string `json:"body"`            // 商品描述信息
	ShowUrl        string `json:"show_url"`        // 商品的展示地址
}

func init() {
	db := Db
	_ = db.AutoMigrate(&Admin{})
	_ = db.AutoMigrate(&LoginLog{})
	_ = db.AutoMigrate(&Orders{})
	_ = db.AutoMigrate(&OrderDetails{})
}
