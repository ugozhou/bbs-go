package common

import "github.com/mlogclub/simple"

var (
	CaptchaError          = simple.NewError(1000, "验证码错误")
	MobileError           = simple.NewError(1001, "电话号码错误")
	InviteCodeError       = simple.NewError(1002, "邀请码错误")
	PaypasswordNotSet     = simple.NewError(1003, "支付密码未设置")
	PaypasswordNotCorrect = simple.NewError(1004, "支付密码错误")
	ProductNotExit        = simple.NewError(1005, "矿机产品不存在")
)
