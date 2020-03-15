package api

import (
	"bbs-go/common"
	"bbs-go/model"
	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple"

	"bbs-go/services"
)

type UserScoreController struct {
	Ctx iris.Context
}

//获取用户矿机资产信息
func (c *UserScoreController) Get() *simple.JsonResult {
	user := services.UserTokenService.GetCurrent(c.Ctx)
	if user == nil {
		return simple.JsonError(simple.ErrorNotLogin)
	}

	userScore := services.UserScoreService.GetByUserId(user.Id)

	if userScore == nil {
		return simple.JsonData(model.UserScore{UserId: user.Id, Score: 0, RewardScore: 0})
	}
	return simple.JsonData(model.UserScore{UserId: user.Id, Score: userScore.RewardScore, RewardScore: userScore.RewardScore})
}

//购买矿机
func (c *UserScoreController) PostBuy(productId, num int, payPassword string) *simple.JsonResult {
	user := services.UserTokenService.GetCurrent(c.Ctx)
	if user == nil {
		return simple.JsonError(simple.ErrorNotLogin)
	}
	if len(user.Paypassword) == 0 {
		return simple.JsonError(common.PaypasswordNotSet)
	}
	if !simple.ValidatePassword(user.Paypassword, payPassword) {
		return simple.JsonError(common.PaypasswordNotCorrect)
	}
	//计算产品金额
	product := services.ProjectService.Get(productId)
	if product == nil {
		return simple.JsonError(common.ProductNotExit)
	}
	if num <= 0 {
		return simple.JsonErrorMsg("参数错误")
	}
	totalPrice := (int64)(num) * product.Price
	totalPower := product.Capacity * num
	//查看余额是否足额。

	// 扣除USDT，增加USDT扣除记录，
	err := services.UserAssertService.BuyMinerCostUserAssert(user.Id, totalPrice)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	//增加算力，增加算力添加记录，给邀请人员增加奖励算力
	err, userScore := services.UserScoreService.IncrementSelfBuyScore(user, totalPower) //一次性购买2T，或购买10T是算几单
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	if userScore == nil {
		return simple.JsonData(model.UserScore{UserId: user.Id, Score: 0, RewardScore: 0, Level: 0})
	}
	return simple.JsonData(model.UserScore{UserId: user.Id, Score: userScore.RewardScore, RewardScore: userScore.RewardScore, Level: userScore.Level})

}
