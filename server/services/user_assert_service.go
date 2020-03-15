package services

import (
	"bbs-go/model"
	"bbs-go/repositories"
	"errors"
	"github.com/mlogclub/simple"
)

var UserAssertService = newUserAssertService()

func newUserAssertService() *userAssertService {
	return &userAssertService{}
}

type userAssertService struct {
}

func (s *userAssertService) Get(id int64) *model.UserAssert {
	return repositories.UserAssertRepository.Get(simple.DB(), id)
}

func (s *userAssertService) Take(where ...interface{}) *model.UserAssert {
	return repositories.UserAssertRepository.Take(simple.DB(), where...)
}

func (s *userAssertService) Find(cnd *simple.SqlCnd) []model.UserAssert {
	return repositories.UserAssertRepository.Find(simple.DB(), cnd)
}

func (s *userAssertService) FindOne(cnd *simple.SqlCnd) *model.UserAssert {
	return repositories.UserAssertRepository.FindOne(simple.DB(), cnd)
}

func (s *userAssertService) FindPageByParams(params *simple.QueryParams) (list []model.UserAssert, paging *simple.Paging) {
	return repositories.UserAssertRepository.FindPageByParams(simple.DB(), params)
}

func (s *userAssertService) FindPageByCnd(cnd *simple.SqlCnd) (list []model.UserAssert, paging *simple.Paging) {
	return repositories.UserAssertRepository.FindPageByCnd(simple.DB(), cnd)
}

func (s *userAssertService) Count(cnd *simple.SqlCnd) int {
	return repositories.UserAssertRepository.Count(simple.DB(), cnd)
}

func (s *userAssertService) Create(t *model.UserAssert) error {
	return repositories.UserAssertRepository.Create(simple.DB(), t)
}

func (s *userAssertService) Update(t *model.UserAssert) error {
	return repositories.UserAssertRepository.Update(simple.DB(), t)
}

func (s *userAssertService) Updates(id int64, columns map[string]interface{}) error {
	return repositories.UserAssertRepository.Updates(simple.DB(), id, columns)
}

func (s *userAssertService) UpdateColumn(id int64, name string, value interface{}) error {
	return repositories.UserAssertRepository.UpdateColumn(simple.DB(), id, name, value)
}

func (s *userAssertService) Delete(id int64) {
	repositories.UserAssertRepository.Delete(simple.DB(), id)
}
func (s *userAssertService) GetByUserId(userId int64) *model.UserAssert {
	return s.FindOne(simple.NewSqlCnd().Eq("user_id", userId))
}

//购买矿机扣除usdt
func (s *userAssertService) BuyMinerCostUserAssert(userId int64, amount int64) error {
	userAssert := s.GetByUserId(userId)
	if userAssert == nil || userAssert.USDT < amount {
		return errors.New("USDT余额不足")
	}
	userAssert.USDT -= amount
	repositories.UserAssertRepository.Update(simple.DB(), userAssert)
	repositories.UserAssertLogRepository.Create(simple.DB(), &model.UserAssertLog{UserId: userId, Amount: amount, Type: model.BuyMiner, Description: "购买矿机"})
	return nil
}

//充值USDT
func (s *userAssertService) RechargeUserAssert(userId int64, amount int64) error {
	userAssert := s.GetByUserId(userId)
	if userAssert == nil {
		return errors.New("没有USDT账户，非法充值，应该先绑定充值地址")
	}
	userAssert.USDT += amount
	repositories.UserAssertRepository.Update(simple.DB(), userAssert)
	repositories.UserAssertLogRepository.Create(simple.DB(), &model.UserAssertLog{UserId: userId, Amount: amount, Type: model.Recharge, Description: "充值USDT"})
	return nil
}
